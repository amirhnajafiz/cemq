package cmd

import (
	"github.com/amirhnajafiz/cemq/pkg/model"

	"github.com/spf13/cobra"
)

// Bench command is used to handle benchmark operations
type Bench struct {
	Input *model.Benchmark
	Cfg   *model.Config
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

}
