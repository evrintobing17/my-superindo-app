package main

import (
	"log"

	"github.com/evrintobing17/my-superindo-app/config"
	"github.com/evrintobing17/my-superindo-app/internal/repository"
	"github.com/evrintobing17/my-superindo-app/pkg/middleware"
	"github.com/gin-gonic/gin"

	authDelivery "github.com/evrintobing17/my-superindo-app/internal/module/auth/delivery/http"
	authRepository "github.com/evrintobing17/my-superindo-app/internal/module/auth/repository"
	authUsecasse "github.com/evrintobing17/my-superindo-app/internal/module/auth/usecase"

	productDelivery "github.com/evrintobing17/my-superindo-app/internal/module/product/delivery/http"
	productRepository "github.com/evrintobing17/my-superindo-app/internal/module/product/repository"
	productUsecase "github.com/evrintobing17/my-superindo-app/internal/module/product/usecase"

	cartDelivery "github.com/evrintobing17/my-superindo-app/internal/module/cart/delivery/http"
	cartRepository "github.com/evrintobing17/my-superindo-app/internal/module/cart/repository"
	cartUsecase "github.com/evrintobing17/my-superindo-app/internal/module/cart/usecase"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database connection
	db, err := repository.NewDatabase(cfg.DB)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	r := gin.New()

	// Initialize Repo
	authRepo := authRepository.NewAuthRepository(db)
	productRepo := productRepository.NewProductRepository(db)
	cartRepo := cartRepository.NewCartRepository(db)

	// Initialize Usecase
	authUC := authUsecasse.NewAuthUsecase(authRepo)
	productUsecase := productUsecase.NewProductUsecase(productRepo)
	cartUsecase := cartUsecase.NewCartUsecase(cartRepo)

	// Initialize Middleware
	middleware := middleware.NewAuthMiddleware(authRepo)

	// Initialize Handler
	authDelivery.NewAuthHandler(r, authUC)
	productDelivery.NewAuthHandler(r, productUsecase, middleware)
	cartDelivery.NewCartHandler(r, cartUsecase, middleware)

	// Initialize API handlers
	// api.SetupRoutes(r, swipeService, premiumService)

	// Start server
	log.Printf("Server running on port %s", cfg.Server.Port)
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
