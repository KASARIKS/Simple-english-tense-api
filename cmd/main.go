package main

import (
	"log"

	"github.com/kasariks/a_nameless_project_yet/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
