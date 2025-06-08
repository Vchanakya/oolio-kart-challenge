package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler aggregates dependencies for HTTP handlers.
type Handler struct {
	db *mongo.Database
}

// NewHandler returns a new Handler.
func NewHandler(db *mongo.Database) *Handler {
	return &Handler{db: db}
}

// ListProducts handles GET /product
func (h *Handler) ListProducts(c echo.Context) error {
	// TODO: Fetch products from MongoDB
	return c.JSON(http.StatusOK, []interface{}{})
}

// GetProduct handles GET /product/:productId
func (h *Handler) GetProduct(c echo.Context) error {
	productID := c.Param("productId")
	_ = productID // TODO: fetch product by ID
	return c.JSON(http.StatusOK, map[string]interface{}{})
}

// PlaceOrder handles POST /order
func (h *Handler) PlaceOrder(c echo.Context) error {
	// TODO: parse request and store order
	return c.JSON(http.StatusOK, map[string]interface{}{})
}
