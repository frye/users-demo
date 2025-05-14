package main

import (
	"log"

	"userprofile-api/api"
)

func main() {
	router := api.SetupRouter()
	
	log.Println("Starting server on :8080")
	router.Run(":8080")
}
