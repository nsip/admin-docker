package main

// Use api and Test
import (
	"fmt"

	"github.com/nsip/admin-docker"
)

func main() {
	doc := admindocker.New()
	doc.UpdateRunning()
	fmt.Printf("%+v\n", doc.Containers)
}
