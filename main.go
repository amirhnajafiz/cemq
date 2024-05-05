package main

import (
	"log"

	"github.com/amirhnajafiz/cemq/internal/mqtt"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

func main() {
	// using cobra-cli for defining commands and subcommands
	// root := cobra.Command{}

	// // execute cobra root command
	// if err := root.Execute(); err != nil {
	// 	panic(err)
	// }

	cli := mqtt.NewCLI(&model.Config{}, false)

	if msg, err := cli.CheckConnection(); err != nil {
		log.Println(err)
	} else {
		log.Println(msg)
	}

	if msg, err := cli.CheckHealth(); err != nil {
		log.Println(err)
	} else {
		log.Println(msg)
	}
}
