package main

import (
	"github.com/JamestarKurbah/go-auth-app/src/routes"
)

func main() {
	router := routes.SetupRoutes()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

// ...existing code...
