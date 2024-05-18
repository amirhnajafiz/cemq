package cmd

import "github.com/spf13/cobra"

// Bench command is used to handle benchmark operations
type Bench struct{}

func (b Bench) Command() *cobra.Command {
	return &cobra.Command{
		Short: "bench",
		Long:  "benchmark",
		Use:   "benchmark the current EMQX cluster",
		Run: func(_ *cobra.Command, _ []string) {
			b.main()
		},
	}
}

func (b Bench) main() {

}
