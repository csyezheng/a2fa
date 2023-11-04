package commands

import (
	"context"
	"fmt"
	"github.com/csyezheng/a2fa/internal/initialize"
	"github.com/spf13/cobra"
	"io"
	"log"
	"log/slog"
	"os"
)

type rootCommand struct {
	Printf  func(format string, v ...interface{})
	Println func(a ...interface{})
	Out     io.Writer

	logger log.Logger

	name     string
	use      string
	commands []Commander
}

func (r *rootCommand) Name() string {
	return r.name
}

func (r *rootCommand) Use() string {
	return r.use
}

func (r *rootCommand) Init(cd *Commandeer) error {
	cmd := cd.CobraCommand
	cmd.Use = "a2fa [flags]"
	cmd.Short = "a2fa generate OTP"
	cmd.Long = `
a2fa is a command line tool for generating and validating one-time password.

See the home page (https://github.com/csyezheng/a2fa/) for more usage.
`
	return nil
}

func (r *rootCommand) Args(ctx context.Context, cd *Commandeer, args []string) error {
	return nil
}

func (r *rootCommand) PreRun(cd, runner *Commandeer) error {
	return nil
}

func (r *rootCommand) Run(ctx context.Context, cd *Commandeer, args []string) error {
	slog.Debug(fmt.Sprintf("a2fa version %q finishing with parameters %q", initialize.Version, os.Args))
	return cd.CobraCommand.Usage()
}

func (r *rootCommand) Commands() []Commander {
	return r.commands
}

func (r *Exec) Execute(ctx context.Context, args []string) (*Commandeer, error) {
	if args == nil {
		// Cobra falls back to os.Args[1:] if args is nil.
		args = []string{}
	}
	r.c.CobraCommand.SetArgs(args)
	cobraCommand, err := r.c.CobraCommand.ExecuteContextC(ctx)
	var cd *Commandeer
	if cobraCommand != nil {
		if err == nil {
			err = checkArgs(cobraCommand, args)
		}

		// Find the commandeer that was executed.
		var find func(*cobra.Command, *Commandeer) *Commandeer
		find = func(what *cobra.Command, in *Commandeer) *Commandeer {
			if in.CobraCommand == what {
				return in
			}
			for _, in2 := range in.commandeers {
				if found := find(what, in2); found != nil {
					return found
				}
			}
			return nil
		}
		cd = find(cobraCommand, r.c)
	}

	return cd, err
}

func Execute(args []string) error {
	// Default GOMAXPROCS to be CPU limit aware, still respecting GOMAXPROCS env.
	x, err := newExec()
	if err != nil {
		return err
	}
	cd, err := x.Execute(context.Background(), args)
	if err != nil {
		if err == errHelp {
			cd.CobraCommand.Help()
			fmt.Println()
			return nil
		}
		if IsCommandError(err) {
			// Print the help, but also return the error to fail the command.
			cd.CobraCommand.Help()
			fmt.Println()
		}
	}
	return err
}

func New(rootCmd Commander) (*Exec, error) {
	rootCd := &Commandeer{
		Command: rootCmd,
	}
	rootCd.Root = rootCd

	// Add all commands recursively.
	var addCommands func(cd *Commandeer, cmd Commander)
	addCommands = func(cd *Commandeer, cmd Commander) {
		cd2 := &Commandeer{
			Root:    rootCd,
			Parent:  cd,
			Command: cmd,
		}
		cd.commandeers = append(cd.commandeers, cd2)
		for _, c := range cmd.Commands() {
			addCommands(cd2, c)
		}

	}

	for _, cmd := range rootCmd.Commands() {
		addCommands(rootCd, cmd)
	}

	if err := rootCd.compile(); err != nil {
		return nil, err
	}

	return &Exec{c: rootCd}, nil

}
