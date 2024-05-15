package cmd

import "github.com/spf13/cobra"

type Bench struct{}

func (b Bench) Command() *cobra.Command {
	return &cobra.Command{
		Short: "bench",
		Long:  "benchmark",
		Use:   "benchmark EMQX cluster",
		Run: func(_ *cobra.Command, _ []string) {
			b.main()
		},
	}
}

func (b Bench) main() {

}
