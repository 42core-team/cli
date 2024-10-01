package docker

import "core-cli/db"

func CleanUP() {
	for _, container := range db.GetContainers() {
		StopContainer(container.ContainerID)
		RemoveContainer(container.ContainerID)
	}

	for _, network := range db.GetNetworks() {
		RemoveNetwork(network.NetworkID)
	}
}
