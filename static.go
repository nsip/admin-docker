package main

import "net/http"
import "fmt"

func main() {
	fmt.Print("Static server started. http://localhost:8097/\n")
	http.ListenAndServe(":8097", http.FileServer(http.Dir("./")))
}
