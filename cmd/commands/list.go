package commands

import (
	"context"
	"fmt"
	"github.com/csyezheng/a2fa/internal/database"
	"github.com/csyezheng/a2fa/internal/initialize"
	"log"
	"log/slog"
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
	return nil
}

func (c *listCommand) PreRun(cd, runner *Commandeer) error {
	c.r = cd.Root.Command.(*rootCommand)
	return nil
}

func (c *listCommand) Run(ctx context.Context, cd *Commandeer, args []string) error {
	initialize.Init()
	if err := c.listAccounts(args); err != nil {
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
		use:  "a2fa list [account name]",
	}
	return listCmd
}

func (c *listCommand) listAccounts(names []string) error {
	db, err := database.LoadDatabase()
	if err != nil {
		log.Fatal("failed to load database: %w", err)
	}
	if err := db.Open(); err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	defer db.Close()

	accounts, err := db.ListAccounts(names)
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		slog.Info("no accounts found!")
	} else {
		for i, account := range accounts {
			fmt.Printf("%d %s %s\n", i, account.Name, account.OTP())
		}
	}
	return nil
}
