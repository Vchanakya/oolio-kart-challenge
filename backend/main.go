package main

import (

	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Create a new Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed to create mongo client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}
	defer client.Disconnect(ctx)

	// Initialize handlers with the database client
	h := NewHandler(client.Database("kart"))

	// Routes
	e.GET("/product", h.ListProducts)
	e.GET("/product/:productId", h.GetProduct)
	e.POST("/order", h.PlaceOrder)

	// Start server
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}

}
