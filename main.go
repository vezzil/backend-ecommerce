package main

import (
	"fmt"
	"log"

	"backend-ecommerce/internal/application/router"
	"backend-ecommerce/internal/infrastructure/cachemanager"
	"backend-ecommerce/internal/infrastructure/config"
	"backend-ecommerce/internal/infrastructure/cronmanager"
	"backend-ecommerce/internal/infrastructure/dbmanager"
	"backend-ecommerce/internal/server/httpserver"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize infrastructure managers
	dbmanager.Init()
	cachemanager.Init()
	cronmanager.Init()

	// Create HTTP server (gin.Engine)
	r := httpserver.New()

	// Get database connection
	db, err := dbmanager.GetDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize JWT
	config.InitJWT()

	// Register application routes
	router.Register(r, db)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	_ = r.Run(addr)
}
