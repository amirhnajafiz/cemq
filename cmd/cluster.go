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
	// cluster health
	root.AddCommand(&cobra.Command{
		Short: "Cluster health check",
		Long:  "Cluster health check uses ping-pong mechanisem to check cluster availability",
		Use:   "health",
	})
	// cluster connection
	root.AddCommand(&cobra.Command{
		Short: "Cluster connection check",
		Long:  "Cluster connection check only checks the EMQX cluster reachability",
		Use:   "connection",
	})
	// finish sub-commands

	return root
}
