package main

import (
	"log"

	"github.com/amirhnajafiz/cemq/cmd"

	"github.com/spf13/cobra"
)

func main() {
	// using cobra-cli for defining commands and subcommands
	root := cobra.Command{}

	// add commands
	root.AddCommand(
		cmd.Config{}.Command(),
		cmd.Bench{}.Command(),
	)

	// execute cobra root command
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
