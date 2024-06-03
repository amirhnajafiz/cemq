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
		debug    bool
		logs     bool
		server   string
		username string
		password string
	)

	// add base flags
	root.PersistentFlags().BoolVar(&debug, "debug", false, "enabling MQTT debug messages")
	root.PersistentFlags().BoolVar(&logs, "logs-enable", true, "disabling or enabling publisher and subscriber logs")
	root.PersistentFlags().StringVar(&server, "server", "", "host address of a EMQX cluster, it will override the config if you set a value to it")
	root.PersistentFlags().StringVar(&username, "username", "", "a username for connecting to a EMQX, only works when the `server` is flag given")
	root.PersistentFlags().StringVar(&password, "password", "", "a password for connecting to a EMQX, only works when the `server` is flag given")

	// load configs
	cfg, err := config.Load(server, username, password)
	if err != nil {
		log.Printf("failed to read context: %v", err)
	}

	// add commands
	root.AddCommand(
		cmd.Config{}.Command(),
		cmd.Load{
			Cfg:   cfg,
			Debug: debug,
			Logs:  logs,
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
