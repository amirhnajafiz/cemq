package cmd

import (
	"log"

	"github.com/amirhnajafiz/cemq/internal/commands/load"
	"github.com/amirhnajafiz/cemq/internal/mqtt"
	"github.com/amirhnajafiz/cemq/pkg/model"

	"github.com/spf13/cobra"
)

// Load command is used to handle load operations
type Load struct {
	Input *model.Load
	Cfg   *model.Config
	Debug bool
}

func (l Load) Command() *cobra.Command {
	root := &cobra.Command{
		Short: "Load on to the current EMQX cluster",
		Long:  "Load on to the current EMQX cluster by giving pubs, subs, and topics",
		Use:   "load",
		Run: func(_ *cobra.Command, _ []string) {
			l.main()
		},
	}

	l.Input = &model.Load{}

	// getting input parameters from flags
	root.Flags().IntVar(&l.Input.Publishers, "pubs", 1, "number of publishers")
	root.Flags().IntVar(&l.Input.Subscribers, "subs", 1, "number of subscriptions")
	root.Flags().IntVar(&l.Input.Topics, "topics", 1, "number of test topics")
	root.Flags().IntVar(&l.Input.PublishIntervals, "interval", 1000, "publish interval in miliseconds")

	return root
}

func (l Load) main() {
	// create a new MQTT connection
	conn := mqtt.NewCLI(l.Cfg, l.Debug)

	// check connection
	if err := conn.Connect(); err != nil {
		log.Fatalf("cannot access the cluster: %v", err)
	}

	// create a new manager
	manager := load.New(conn)

	// start the load generate
	if err := manager.Generate(l.Input); err != nil {
		log.Fatalf("failed to start the load: %v", err)
	}
}
