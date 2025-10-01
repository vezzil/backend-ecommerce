package main

import (
	"fmt"
	"backend-ecommerce/internal/infrastructure/config"
	"backend-ecommerce/internal/infrastructure/dbmanager"
	"backend-ecommerce/internal/infrastructure/cachemanager"
	"backend-ecommerce/internal/infrastructure/cronmanager"
	"backend-ecommerce/internal/server/httpserver"
	"backend-ecommerce/internal/application/router"
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

	// Register application routes
	router.Register(r)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	_ = r.Run(addr)
}
