package main

import (
	"gladius-edge-daemon/init/service-manager"
	"gladius-edge-daemon/internal/network-daemon"
)

// Main entry-point for the service
func main() {
	// Define some variables
	name, displayName, description :=
		"GladiusEdgeDaemon",
		"Gladius Network (Edge) Daemon",
		"Gladius Network (Edge) Daemon"

	// Run the function "run" as a service
	manager.RunService(name, displayName, description, networkd.Run)
}
