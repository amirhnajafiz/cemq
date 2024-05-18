package cmd

import "github.com/spf13/cobra"

// Bench command is used to handle benchmark operations
type Bench struct{}

func (b Bench) Command() *cobra.Command {
	return &cobra.Command{
		Short: "Benchmark the current EMQX cluster",
		Long:  "Benchmark the current EMQX cluster by giving pubs, subs, and topics",
		Use:   "bench",
		Run: func(_ *cobra.Command, _ []string) {
			b.main()
		},
	}
}

func (b Bench) main() {

}
