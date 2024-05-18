package main

import (
	"log"

	"github.com/amirhnajafiz/cemq/cmd"
	"github.com/amirhnajafiz/cemq/internal/config"

	"github.com/spf13/cobra"
)

func main() {
	// using cobra-cli for defining commands and subcommands
	root := cobra.Command{}

	var (
		debug bool
	)

	// add base flags
	root.PersistentFlags().BoolVar(&debug, "debug", false, "enabling MQTT debug messages")

	// load configs
	cfg, err := config.Load()
	if err != nil {
		log.Printf("failed to read context: %v", err)
	}

	// add commands
	root.AddCommand(
		cmd.Config{}.Command(),
		cmd.Bench{
			Cfg:   cfg,
			Debug: debug,
		}.Command(),
		cmd.Cluster{
			Cfg:   cfg,
			Debug: debug,
		}.Command(),
	)

	// execute cobra root command
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
