package admindocker

type AdminDockerImage struct {
	ID      string
	RepoTag string
}

type AdminDockerContainer struct {
	ID    string
	Image string
	Ports []uint16
}

type AdminDockerStore struct {
	Containers []AdminDockerContainer
	Images     []AdminDockerImage
}

// Create a new SIF client, which requires an environment URL
// application key and password. Returns the new object used for requets.
func New() *AdminDockerStore {
	s := &AdminDockerStore{}
	return s
}
