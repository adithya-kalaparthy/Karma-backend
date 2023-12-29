package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/karma/karma-backend/internal/config"
	"github.com/karma/karma-backend/pkg/api/routes"
)

func main() {
	jsonData, err := config.LoadData()

	fmt.Print(jsonData)

	if err != nil {
		log.Fatalf("Error with config %v", err)
		return
	}

	router := gin.Default()
	taskApiGroup := router.Group("/api")

	routes.SetupTaskRoutes(taskApiGroup)

	port := ":8080"

	log.Printf("Server listening on %s", port)

	errRunningRouter := router.Run(port)

	if errRunningRouter != nil {
		log.Fatalf("Router cannot be run %v", errRunningRouter)
		return
	}
}