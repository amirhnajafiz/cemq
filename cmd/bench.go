package cmd

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/cemq/internal/commands/bench"
	"github.com/amirhnajafiz/cemq/internal/mqtt"
	"github.com/amirhnajafiz/cemq/pkg/model"

	"github.com/spf13/cobra"
)

// Bench command is used to handle benchmark operations
type Bench struct {
	Input *model.Benchmark
	Cfg   *model.Config
	Debug bool
}

func (b Bench) Command() *cobra.Command {
	root := &cobra.Command{
		Short: "Benchmark the current EMQX cluster",
		Long:  "Benchmark the current EMQX cluster by N messages and giving pubs, subs, and topic",
		Use:   "bench",
		Run: func(_ *cobra.Command, _ []string) {
			b.main()
		},
	}

	b.Input = &model.Benchmark{}

	// input parameters
	root.Flags().IntVar(&b.Input.Pubs, "pubs", 0, "number of publishers")
	root.Flags().IntVar(&b.Input.Subs, "subs", 0, "number of subscriptions")
	root.Flags().IntVar(&b.Input.Messages, "messages", 0, "number of messages")
	root.Flags().StringVar(&b.Input.Topic, "topic", "test-topic", "topic name")

	return root
}

func (b Bench) main() {
	// create a new MQTT connection
	conn := mqtt.NewCLI(b.Cfg, b.Debug)

	// check connection
	if err := conn.Connect(); err != nil {
		log.Fatalf("cannot access the cluster: %v", err)
	}

	// creating a manager
	manager := bench.New(conn)

	// run benchmark
	upshot := manager.Benchmark(b.Input)

	// display output
	fmt.Printf("benchmakr %s: topic %s\n", upshot.ID, upshot.Topic)
	for _, item := range upshot.Pubs {
		fmt.Printf("publish: %f mps, %f kbyte/s\n", item.WriteMps, item.WriteThroughput)
	}
	for _, item := range upshot.Subs {
		fmt.Printf("subscribe: %f mps, %f kbyte/s\n", item.ReadMps, item.ReadThroughput)
	}
}
