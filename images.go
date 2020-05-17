package admindocker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func (s *AdminDockerStore) UpdateImages() bool {
	// ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	/*
		ID:sha256:9bc90699d78e706164ddadc4eebd3482c7fc2cf6e27444aeeb8720aa88148f74 Labels:map[] ParentID:sha256:9a4ef247f1a2ba92904ae551423cb3d66040f56d6140724ee4a39a680274db9a RepoDigests:[] RepoTags:[nsip/hitsproxy:latest] SharedSize:-1 Size:110000493 VirtualSize:110000493} {Containers:-1 Created:1589714232
	*/

	b, err := json.MarshalIndent(images, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
	fmt.Printf("%+v\n", images)

	s.Images = nil
	for _, image := range images {

		for _, repotag := range image.RepoTags {
			s.Images = append(s.Images, AdminDockerImage{
				ID:      string(image.ID),
				RepoTag: string(repotag),
			})
		}
	}

	return true
}
