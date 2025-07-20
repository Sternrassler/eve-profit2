package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"eve-profit2/internal/api/handlers"
	"eve-profit2/internal/config"
	"eve-profit2/internal/repository"
	"eve-profit2/internal/service"
	"eve-profit2/pkg/esi"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting EVE Profit Calculator Backend...")

	// Load configuration
	cfg := config.Load()

	// Initialize SDE repository
	sdeRepo, err := repository.NewSDERepository(cfg.SDEDatabasePath)
	if err != nil {
		fmt.Printf("Failed to initialize SDE repository: %v\n", err)
		os.Exit(1)
	}
	defer sdeRepo.Close()

	// Initialize ESI client
	esiClient := esi.NewESIClient(
		esi.WithBaseURL(cfg.ESIBaseURL),
		esi.WithRateLimit(cfg.ESIRateLimit),
		esi.WithRetryAttempts(3),
	)

	// Initialize services
	itemService := service.NewItemService(sdeRepo, nil)
	marketService := service.NewMarketService(esiClient)

	// Setup Gin router
	if !cfg.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Basic CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", cfg.CORSOrigin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Basic routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "EVE Profit Calculator API",
			"version": "2.0.0",
			"status":  "running",
			"config": gin.H{
				"port":         cfg.ServerPort,
				"esi_base":     cfg.ESIBaseURL,
				"rate_limit":   cfg.ESIRateLimit,
				"debug_mode":   cfg.DebugMode,
				"client_id":    cfg.ESIClientID,
				"callback_url": cfg.ESICallbackURL,
				"scopes":       cfg.ESIScopes,
			},
		})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "healthy",
				"time":   time.Now(),
			})
		})

		// Test ESI connection
		api.GET("/esi/test", func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			// Test with Tritanium in The Forge
			orders, err := esiClient.GetMarketOrders(ctx, 10000002, 34)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"esi_status": "connected",
				"test_data":  fmt.Sprintf("Retrieved %d market orders", len(orders)),
				"sample":     orders[:min(3, len(orders))], // Show first 3 orders
			})
		})

		// Test SDE connection
		api.GET("/sde/test", func(c *gin.Context) {
			item, err := itemService.GetItemByID(34) // Tritanium
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"sde_status": "connected",
				"test_item":  item,
			})
		})

		// Auth placeholder
		api.GET("/auth/login", func(c *gin.Context) {
			// Placeholder usage to prevent unused variable warning
			_ = marketService
			c.JSON(http.StatusOK, gin.H{
				"message":   "EVE SSO authentication endpoint",
				"client_id": cfg.ESIClientID,
				"callback":  cfg.ESICallbackURL,
				"scopes":    cfg.ESIScopes,
				"status":    "configured",
			})
		})

		// Items API endpoints
		itemsHandler := handlers.NewItemHandler(itemService)
		api.GET("/items/:item_id", itemsHandler.GetItemDetails)
		api.GET("/items/search", itemsHandler.SearchItems)
	}

	// Start server
	srv := &http.Server{
		Addr:         cfg.GetServerAddress(),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		fmt.Printf("🚀 Server starting on port %s\n", cfg.ServerPort)
		fmt.Printf("📡 EVE ESI Client ID: %s\n", cfg.ESIClientID)
		fmt.Printf("🔗 Callback URL: %s\n", cfg.ESICallbackURL)
		fmt.Printf("🌐 API available at: http://localhost:%s\n", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to start server: %v\n", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Server exited")
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
