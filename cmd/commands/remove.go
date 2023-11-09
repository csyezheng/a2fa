package commands

import (
	"context"
	"fmt"
	"github.com/csyezheng/a2fa/internal/database"
	"github.com/csyezheng/a2fa/internal/initialize"
	"github.com/spf13/cobra"
	"log"
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
	if err := cobra.MinimumNArgs(1)(cd.CobraCommand, args); err != nil {
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
	if err := cobra.MinimumNArgs(1)(cd.CobraCommand, args); err != nil {
		return err
	}
	if err := c.removeAccounts(args); err != nil {
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
		use:  "a2fa remove <account name> <account name>...",
	}
	return removeCmd
}

func (c *removeCommand) removeAccounts(names []string) error {
	db, err := database.LoadDatabase()
	if err != nil {
		log.Fatal("failed to load database: %w", err)
	}
	if err := db.Open(); err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	defer db.Close()

	if err := db.RemoveAccounts(names); err != nil {
		return err
	}
	return nil
}
