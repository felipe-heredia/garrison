package main

import (
	"garrison/router"
)

func main() {
	router := router.SetupRouter()
	router.Run(":8080")
}
