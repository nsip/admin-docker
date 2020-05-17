package admindocker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func (s *AdminDockerStore) UpdateRunning() bool {
	// ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	/*
		ID:30c
		Image:nsip/hitsproxy
		Command:/go/bin/hitsproxy
		Created:1589669100
		Ports:[{IP:0.0.0.0 PrivatePort:8089 PublicPort:8089 Type:tcp}]
		SizeRw:0 SizeRootFs:0 Labels:map[] State:running
		Status:Up 14 minutes HostConfig:{NetworkMode:default}
		NetworkSettings:0xc0000105b0 Mounts:[]}
	*/

	s.Containers = nil
	for _, container := range containers {
		ports := make([]uint16, 0)

		// XXX map?
		for _, port := range container.Ports {
			ports = append(ports, port.PublicPort)
		}
		s.Containers = append(s.Containers, AdminDockerContainer{
			ID:    string(container.ID),
			Image: container.Image,
			Ports: ports,
		})
	}

	return true
}
