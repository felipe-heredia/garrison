package main

import (
	database "garrison/internal/database/postgres"

	"garrison/router"
)

func main() {
	router := router.SetupRouter()

	database.Initialize()

	router.Run(":8080")
}
