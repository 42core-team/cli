package docker

import "core-cli/db"

func CleanUP() {
	for _, container := range db.GetContainers() {
		StopContainer(container.ContainerID)
		RemoveContainer(container.ContainerID)
		db.DeleteContainer(container.ContainerID)
	}

	for _, network := range db.GetNetworks() {
		RemoveNetwork(network.NetworkID)
		db.DeleteNetwork(network.NetworkID)
	}

	db.DeleteAllContainers()
	db.DeleteAllNetworks()
}
