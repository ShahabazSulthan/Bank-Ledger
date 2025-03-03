package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ShahabazSulthan/Api-Gateway/pkg/config"
	"github.com/ShahabazSulthan/Api-Gateway/pkg/di"
)

const serverID = "SERVER-8000"

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[%s] Failed to load config: %v", serverID, err)
	}

	router := gin.Default()

	if err := di.InitBankClient(router, cfg); err != nil {
		log.Fatalf("[%s] Failed to initialize Bank Client: %v", serverID, err)
	}

	fmt.Printf("Server running on port %s\n", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
