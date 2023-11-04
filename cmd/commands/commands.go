package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// Commander is the interface that must be implemented by all commands.
type Commander interface {
	Name() string

	Use() string

	// Init is called when the cobra command is created.
	// This is where the flags, short and long description etc. can be added.
	Init(*Commandeer) error

	// Args the command args
	Args(ctx context.Context, cd *Commandeer, args []string) error

	// PreRun called on all ancestors and the executing command itself, before execution, starting from the root.
	// This is the place to evaluate flags and set up the this Commandeer.
	// The runner Commandeer holds the currently running command, which will be PreRun last.
	PreRun(this, runner *Commandeer) error

	// Run the command execution.
	Run(ctx context.Context, cd *Commandeer, args []string) error

	// Commands returns the sub commands, if any.
	Commands() []Commander
}

// Commandeer holds the state of a command and its subcommands.
type Commandeer struct {
	Command      Commander
	CobraCommand *cobra.Command

	Root        *Commandeer
	Parent      *Commandeer
	commandeers []*Commandeer
}

type Exec struct {
	c *Commandeer
}

func checkArgs(cmd *cobra.Command, args []string) error {
	// no subcommand, always take args.
	if !cmd.HasSubCommands() {
		return nil
	}

	var commandName string
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			break
		}
		commandName = arg
	}

	if commandName == "" || cmd.Name() == commandName {
		return nil
	}

	// Also check the aliases.
	if cmd.HasAlias(commandName) {
		return nil
	}

	return fmt.Errorf("unknown command %q for %q%s", args[0], cmd.CommandPath(), findSuggestions(cmd, commandName))
}

func (c *Commandeer) init() error {

	// Collect all ancestors including self.
	var ancestors []*Commandeer
	{
		cd := c
		for cd != nil {
			ancestors = append(ancestors, cd)
			cd = cd.Parent
		}
	}

	// Init all of them starting from the root.
	for i := len(ancestors) - 1; i >= 0; i-- {
		cd := ancestors[i]
		if err := cd.Command.PreRun(cd, c); err != nil {
			return err
		}
	}

	return nil

}

func (c *Commandeer) compile() error {
	useCommandFlagsArgs := "[command] [flags]"
	if len(c.commandeers) == 0 {
		useCommandFlagsArgs = "[flags] [args]"
	}
	c.CobraCommand = &cobra.Command{
		Use: fmt.Sprintf("%s %s", c.Command.Name(), useCommandFlagsArgs),
		Args: func(cmd *cobra.Command, args []string) error {
			if err := c.Command.Args(cmd.Context(), c, args); err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := c.Command.Run(cmd.Context(), c, args); err != nil {
				return err
			}
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return c.init()
		},
		SilenceErrors:              false,
		SilenceUsage:               false,
		SuggestionsMinimumDistance: 2,
	}

	// This is where the flags, short and long description etc. are added
	if err := c.Command.Init(c); err != nil {
		return err
	}

	// Add commands recursively.
	for _, cc := range c.commandeers {
		if err := cc.compile(); err != nil {
			return err
		}
		c.CobraCommand.AddCommand(cc.CobraCommand)
	}

	return nil
}

//var (
//	// Used for flags.
//	cfgFile string
//
//	rootCmd = &cobra.Command{
//		Use:   "a2fa",
//		Short: "A command line tool for generating and validating OTPs",
//		Long:  `a2fa is command line tool for generating and validating OTPs.`,
//	}
//)
//
//func Execute() error {
//	return rootCmd.Execute()
//}
//
//func initialize() {
//	cobra.OnInitialize(initConfig)
//
//	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./configs/config.toml)")
//
//	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
//
//	viper.SetDefault("author", "Ye Zheng <csyezheng@gmail.com>")
//	viper.SetDefault("license", "apache")
//
//	rootCmd.AddCommand(newListCommand())
//	rootCmd.AddCommand(newAddCommand())
//	rootCmd.AddCommand(newRemoveCommand())
//}
//
//func initConfig() {
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		cwd, err := os.Getwd()
//		cfgdir := path.Join(cwd, "configs")
//		cobra.CheckErr(err)
//
//		// Search config in home directory with name ".cobra" (without extension).
//		viper.AddConfigPath(cfgdir)
//		viper.SetConfigType("toml")
//		viper.SetConfigName(".toml")
//	}
//
//	viper.AutomaticEnv()
//
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Println("Using config file:", viper.ConfigFileUsed())
//	}
//}

func findSuggestions(cmd *cobra.Command, arg string) string {
	if cmd.DisableSuggestions {
		return ""
	}
	suggestionsString := ""
	if suggestions := cmd.SuggestionsFor(arg); len(suggestions) > 0 {
		suggestionsString += "\n\nDid you mean this?\n"
		for _, s := range suggestions {
			suggestionsString += fmt.Sprintf("\t%v\n", s)
		}
	}
	return suggestionsString
}

func newExec() (*Exec, error) {
	rootCmd := &rootCommand{
		name: "a2fa",
		use:  "a2fa <subcommand> [flags] [args]",
		commands: []Commander{
			newVersionCommand(),
			newGenerateCommand(),
			newAddCommand(),
			newRemoveCommand(),
			newUpdateCommand(),
			newListCommand(),
		},
	}
	return New(rootCmd)
}
