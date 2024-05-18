package cmd

import (
	"fmt"

	"github.com/amirhnajafiz/cemq/internal/commands/config"

	"github.com/spf13/cobra"
)

// Config command is used to handle config operations
type Config struct{}

func (c Config) Command() *cobra.Command {
	root := &cobra.Command{
		Use:   "config command handles config operations",
		Short: "config",
		Long:  "config",
	}

	manager := config.New()

	// add sub-commands
	// config select {context}
	root.AddCommand(&cobra.Command{
		Short: "select",
		Long:  "select",
		Use:   "select is used to change the config context",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println(manager.Select(args[0]))
		},
	})
	root.AddCommand(&cobra.Command{
		Short: "info",
		Long:  "info",
		Use:   "info is used to get current context info",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println(manager.Info())
		},
	})
	root.AddCommand(&cobra.Command{
		Short: "list",
		Long:  "list",
		Use:   "list is used to print the list of current configs",
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
