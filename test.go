package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	// ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", containers)

	/*

[

{
	ID:30c3170c2d3b853b2295db2eb32f616327ce97d9bde3373b9fdb7750d350470e Names:[/nice_sinoussi] 
	Image:nsip/hitsproxy ImageID:sha256:a48c948fb7b1702ffa6a3c5367245120a31b8271b91cee97e31a00df50d610d0 
	Command:/go/bin/hitsproxy 
	Created:1589669100 
	Ports:[{IP:0.0.0.0 PrivatePort:8089 PublicPort:8089 Type:tcp}] 
	SizeRw:0 SizeRootFs:0 Labels:map[] State:running 
	Status:Up 14 minutes HostConfig:{NetworkMode:default} 
	NetworkSettings:0xc0000105b0 Mounts:[]} 

{ID:8de4ae22ec5ee32303c63c05cfb56292c375ff2594efcacd5fe608740e9dfa4a Names:[/interesting_aryabhata] Image:nats-streaming ImageID:sha256:411737a82b95162c3621ea19f4edd42fc7c9e58e99cce32f343de607bea1b89b Command:/nats-streaming-server -p 4223 -m 8223 Created:1589409069 Ports:[{IP:0.0.0.0 PrivatePort:4223 PublicPort:4223 Type:tcp} {IP: PrivatePort:8222 PublicPort:0 Type:tcp} {IP:0.0.0.0 PrivatePort:8223 PublicPort:8223 Type:tcp} {IP: PrivatePort:4222 PublicPort:0 Type:tcp}] SizeRw:0 SizeRootFs:0 Labels:map[] State:running Status:Up 3 days HostConfig:{NetworkMode:default} NetworkSettings:0xc0000105d0 Mounts:[]}

]

Container=30c3170c2d3b853b2295db2eb32f616327ce97d9bde3373b9fdb7750d350470e,Image=nsip/hitsproxy,Ports=[{0.0.0.0 %!s(uint16=8089) %!s(uint16=8089) tcp}]

Container=8de4ae22ec5ee32303c63c05cfb56292c375ff2594efcacd5fe608740e9dfa4a,Image=nats-streaming,Ports=[{0.0.0.0 %!s(uint16=4223) %!s(uint16=4223) tcp} { %!s(uint16=8222) %!s(uint16=0) tcp} {0.0.0.0 %!s(uint16=8223) %!s(uint16=8223) tcp} { %!s(uint16=4222) %!s(uint16=0) tcp}]


	*/


	for _, container := range containers {
		fmt.Printf("Container=%s,Image=%s,Ports=%s\n", container.ID, container.Image, container.Ports)
	}
}
