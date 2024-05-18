package cmd

import (
	"fmt"

	"github.com/amirhnajafiz/cemq/internal/commands/config"

	"github.com/spf13/cobra"
)

// Config command is used to handle config operations
type Config struct{}

func (c Config) Command() *cobra.Command {
	// creating root command
	root := &cobra.Command{
		Short: "Config handles config operations",
		Long:  "Config command hanldes config operations for list, view, and change config files",
		Use:   "config",
	}

	// creating the manager
	manager := config.New()

	// add sub-commands
	// config select {context}
	root.AddCommand(&cobra.Command{
		Short: "Select is used to change the config context",
		Long:  "Select is used to change the config context by setting the context name without json extention",
		Use:   "select [context]",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println(manager.Select(args[0]))
		},
	})
	// config info
	root.AddCommand(&cobra.Command{
		Short: "Info is used to get current context info",
		Long:  "Info reads the current context and prints its data as json output",
		Use:   "info",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println(manager.Info())
		},
	})
	// config list
	root.AddCommand(&cobra.Command{
		Short: "List is used to print the list of current configs",
		Long:  "List is used to print the list of current configs",
		Use:   "list",
		Run: func(_ *cobra.Command, args []string) {
			list := manager.List()
			for _, item := range list {
				fmt.Println(item)
			}
		},
	})
	// finish sub-commands

	return root
}
