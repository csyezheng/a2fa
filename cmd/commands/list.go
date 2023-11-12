package commands

import (
	"context"
	"fmt"
	"github.com/csyezheng/a2fa/internal/database"
	"github.com/csyezheng/a2fa/internal/initialize"
	"github.com/csyezheng/a2fa/internal/models"
	"github.com/spf13/cobra"
	"log"
	"sync"
)

type listCommand struct {
	r        *rootCommand
	name     string
	use      string
	commands []Commander
}

func (c *listCommand) Name() string {
	return c.name
}

func (c *listCommand) Use() string {
	return c.use
}

func (c *listCommand) Init(cd *Commandeer) error {
	cmd := cd.CobraCommand
	cmd.Short = "List all added accounts and password code"
	cmd.Long = "List all added accounts and password code"
	return nil
}

func (c *listCommand) Args(ctx context.Context, cd *Commandeer, args []string) error {
	if err := cobra.RangeArgs(0, 2)(cd.CobraCommand, args); err != nil {
		return err
	}
	return nil
}

func (c *listCommand) PreRun(cd, runner *Commandeer) error {
	c.r = cd.Root.Command.(*rootCommand)
	return nil
}

func (c *listCommand) Run(ctx context.Context, cd *Commandeer, args []string) error {
	initialize.Init()
	var accountName, userName string
	if len(args) == 1 {
		if pairs := strings.SplitN(args[0], ":", 2); len(pairs) == 2 {
			accountName = pairs[0]
			userName = pairs[1]
		} else {
			accountName = args[0]
		}
	} else if len(args) == 2 {
		accountName = args[0]
		userName = args[1]
	}
	if err := c.listAccounts(accountName, userName); err != nil {
		log.Fatal(err)
	}
	return nil
}

func (c *listCommand) Commands() []Commander {
	return c.commands
}

func newListCommand() *listCommand {
	listCmd := &listCommand{
		name: "list",
		use:  "list [account name]",
	}
	return listCmd
}

func (c *listCommand) listAccounts(accountName string, userName string) error {
	db, err := database.LoadDatabase()
	if err != nil {
		log.Fatal("failed to load database: %w", err)
	}
	if err := db.Open(); err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	defer db.Close()

	accounts, err := db.ListAccounts(accountName, userName)
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		fmt.Println("no accounts found!")
	} else {
		var wg sync.WaitGroup
		for _, account := range accounts {
			wg.Add(1)
			go func(account models.Account) {
				defer wg.Done()
				code, err := account.OTP()
				if err != nil {
					log.Printf("%s %s generate code error%s\n", account.AccountName, account.Username, err)
				} else {
					log.Printf("%s %s %s\n", account.AccountName, account.Username, code)
				}
			}(account)
		}
		wg.Wait()
	}
	return nil
}
