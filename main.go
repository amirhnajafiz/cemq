package main

import "github.com/spf13/cobra"

func main() {
	// using cobra-cli for defining commands and subcommands
	root := cobra.Command{}

	// execute cobra root command
	if err := root.Execute(); err != nil {
		panic(err)
	}
}
