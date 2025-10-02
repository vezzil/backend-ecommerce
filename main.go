package main

import (
	"fmt"

	"backend-ecommerce/internal/application/router"
	"backend-ecommerce/internal/infrastructure/cachemanager"
	"backend-ecommerce/internal/infrastructure/config"
	"backend-ecommerce/internal/infrastructure/cronmanager"
	"backend-ecommerce/internal/infrastructure/dbmanager"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize infrastructure managers
	dbmanager.Init()
	cachemanager.Init()
	cronmanager.Init()

	// Create Gin router with default middleware
	r := gin.Default()

	// Apply CORS middleware from router_config
	r.Use(config.CORSMiddleware())

	// Get database connection
	db := dbmanager.GetDB()


	// Initialize JWT
	config.InitJWT()

	// Register application routes
	router.Register(r, db)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	_ = r.Run(addr)
}
