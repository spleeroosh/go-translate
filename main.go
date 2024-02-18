package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spleeroosh/go-translate/config"
	"github.com/spleeroosh/go-translate/database"
	"github.com/spleeroosh/go-translate/handler"
)

const (
	maxDBConnectionAttempts = 10
	dbConnectionRetryDelay  = 5 * time.Second
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load("/app/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Load configuration from environment variables
	cfg := config.NewConfig()
	fmt.Println(cfg)

	// Initialize database connection with retry logic
	var db *database.DB
	var err error
	for attempt := 1; attempt <= maxDBConnectionAttempts; attempt++ {
		db, err = database.Connect(cfg.DatabaseURL())
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", attempt, maxDBConnectionAttempts, err)
		time.Sleep(dbConnectionRetryDelay)
	}
	if err != nil {
		log.Fatalf("Exceeded maximum number of database connection attempts: %v", err)
	}
	defer db.Close()

	// Initialize translation store
	translationStore := database.NewTranslationStore(db.DB)

	fmt.Println(translationStore)

	// Inject translation store into handlers
	homeHandler := handler.HomeHandler{}
	authHandler := handler.AuthHandler{}
	translationsHandler := handler.TranslationsHandler{}

	app.Use(withUser)
	app.GET("/", homeHandler.HandleHomeShow)
	app.GET("/auth", authHandler.HandleAuthShow)
	app.GET("/translations", translationsHandler.HandleTranslationsShow)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Logger.Fatal(app.Start(":" + port))
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "rogovin.pavel@mail.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
