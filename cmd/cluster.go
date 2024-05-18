package cmd

import (
	"github.com/spf13/cobra"
)

// Cluster command is used to handle cluster check and connection methods
type Cluster struct{}

func (c Cluster) Command() *cobra.Command {
	// creating root command
	root := &cobra.Command{
		Short: "Cluster handles cluster operations",
		Long:  "Config command hanldes cluster operations for connection check and health",
		Use:   "cluster",
	}

	// add sub-commands

	// finish sub-commands

	return root
}
