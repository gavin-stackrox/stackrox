package main

import (
	"github.com/spf13/cobra"
	"github.com/stackrox/rox/cmd/deploy/central"
	"github.com/stackrox/rox/generated/api/v1"
)

func dockerBasedOrchestrator(shortName, longName string, cluster v1.ClusterType) *cobra.Command {
	swarmConfig := new(central.SwarmConfig)

	c := orchestratorCommand(shortName, longName)
	c.PersistentPreRun = func(*cobra.Command, []string) {
		cfg.SwarmConfig = swarmConfig
		cfg.ClusterType = cluster
	}
	c.AddCommand(externalVolume())
	c.AddCommand(hostPathVolume(cluster))
	c.AddCommand(noVolume())

	// Adds swarm specific flags
	c.PersistentFlags().StringVarP(&swarmConfig.ClairifyImage, "clairify-image", "", "stackrox.io/"+clairifyImage, "Clairify image to use")
	c.PersistentFlags().StringVarP(&swarmConfig.MainImage, "main-image", "i", "stackrox.io/"+mainImage, "Tmage to use")
	c.PersistentFlags().StringVarP(&swarmConfig.NetworkMode, "mode", "m", "ingress", "network mode to use (ingress or host)")
	c.PersistentFlags().IntVarP(&swarmConfig.PublicPort, "port", "p", 443, "public port to expose")
	return c
}
