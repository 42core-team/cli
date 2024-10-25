package docker

import "core-cli/db"

func CleanUP() {
	for _, container := range db.GetContainers() {
		StopRmContainer(container.ContainerID)
	}

	for _, network := range db.GetNetworks() {
		StopRmNetwork(network.NetworkID)
	}

	db.DeleteAllContainers()
	db.DeleteAllNetworks()
}

func StopRmContainer(containerID string) {
	StopContainer(containerID)
	RemoveContainer(containerID)
	db.DeleteContainer(containerID)
}

func StopRmNetwork(networkID string) {
	RemoveNetwork(networkID)
	db.DeleteNetwork(networkID)
}
