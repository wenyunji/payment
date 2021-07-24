package main

import (
	"log"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("payment"),
		micro.Version("latest"),
	)

	// Register handler

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
