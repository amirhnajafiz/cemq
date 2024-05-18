package cmd

import (
	"fmt"

	"github.com/amirhnajafiz/cemq/internal/commands/cluster"
	"github.com/amirhnajafiz/cemq/internal/mqtt"

	"github.com/spf13/cobra"
)

// Cluster command is used to handle cluster check and connection methods
type Cluster struct {
	conn mqtt.CLI
}

func (c Cluster) Command() *cobra.Command {
	// creating root command
	root := &cobra.Command{
		Short: "Cluster handles cluster operations",
		Long:  "Config command hanldes cluster operations for connection check and health",
		Use:   "cluster",
	}

	// creating the manager
	manager := cluster.New(c.conn)

	// add sub-commands
	// cluster health
	root.AddCommand(&cobra.Command{
		Short: "Cluster health check",
		Long:  "Cluster health check uses ping-pong mechanisem to check cluster availability",
		Use:   "health",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(manager.CheckHealth())
		},
	})
	// cluster connection
	root.AddCommand(&cobra.Command{
		Short: "Cluster connection check",
		Long:  "Cluster connection check only checks the EMQX cluster reachability",
		Use:   "connection",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(manager.CheckConnection())
		},
	})
	// finish sub-commands

	return root
}
