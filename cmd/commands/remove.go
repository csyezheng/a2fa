package commands

import (
	"context"
	"fmt"
	"github.com/csyezheng/a2fa/internal/database"
	"github.com/csyezheng/a2fa/internal/initialize"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type removeCommand struct {
	r *rootCommand

	name     string
	use      string
	commands []Commander
}

func (c *removeCommand) Name() string {
	return c.name
}

func (c *removeCommand) Use() string {
	return c.use
}

func (c *removeCommand) Init(cd *Commandeer) error {
	cmd := cd.CobraCommand
	cmd.Short = "Remove account and its secret key"
	cmd.Long = "Remove account and its secret key"
	return nil
}

func (c *removeCommand) Args(ctx context.Context, cd *Commandeer, args []string) error {
	if err := cobra.RangeArgs(1, 2)(cd.CobraCommand, args); err != nil {
		return err
	}
	return nil
}

func (c *removeCommand) PreRun(cd, runner *Commandeer) error {
	c.r = cd.Root.Command.(*rootCommand)
	return nil
}

func (c *removeCommand) Run(ctx context.Context, cd *Commandeer, args []string) error {
	initialize.Init()
	accountName, userName := args[0], ""
	if len(args) == 1 {
		if pairs := strings.SplitN(args[0], ":", 2); len(pairs) == 2 {
			accountName = pairs[0]
			userName = pairs[1]
		}
	} else {
		accountName = args[0]
		userName = args[1]
	}
	if err := c.removeAccount(accountName, userName); err != nil {
		log.Fatal(err)
	}
	fmt.Println("accounts deleted successfully")
	return nil
}

func (c *removeCommand) Commands() []Commander {
	return c.commands
}

func newRemoveCommand() *removeCommand {
	removeCmd := &removeCommand{
		name: "remove",
		use:  "remove <account name> [user name]",
	}
	return removeCmd
}

func (c *removeCommand) removeAccount(accountName string, userName string) error {
	db, err := database.LoadDatabase()
	if err != nil {
		log.Fatal("failed to load database: %w", err)
	}
	if err := db.Open(); err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	defer db.Close()

	if err := db.RemoveAccount(accountName, userName); err != nil {
		return err
	}
	return nil
}
