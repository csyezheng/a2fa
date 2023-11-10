package commands

import (
	"context"
	"fmt"
	"github.com/csyezheng/a2fa/internal/initialize"
)

type versionCommand struct {
	r        *rootCommand
	name     string
	use      string
	commands []Commander
}

func (c *versionCommand) Name() string {
	return c.name
}

func (c *versionCommand) Use() string {
	return c.use
}

func (c *versionCommand) Init(cd *Commandeer) error {
	cmd := cd.CobraCommand
	cmd.Short = "show version"
	cmd.Long = "show version"
	return nil
}

func (c *versionCommand) Args(ctx context.Context, cd *Commandeer, args []string) error {
	return nil
}

func (c *versionCommand) PreRun(cd, runner *Commandeer) error {
	c.r = cd.Root.Command.(*rootCommand)
	return nil
}

func (c *versionCommand) Run(ctx context.Context, cd *Commandeer, args []string) error {
	c.ShowVersion()
	return nil
}

func (c *versionCommand) Commands() []Commander {
	return c.commands
}

func newVersionCommand() *versionCommand {
	versionCmd := &versionCommand{
		name: "version",
		use:  "version",
	}
	return versionCmd
}

func (c *versionCommand) ShowVersion() {
	fmt.Printf("a2fa %s\n", initialize.Version)
}
