package repository

import (
	"context"
	"testing"

	"github.com/Vchanakya/oolio-kart-challenge/backend/db"
	"github.com/Vchanakya/oolio-kart-challenge/backend/internal/models"
	"github.com/Vchanakya/oolio-kart-challenge/backend/internal/seed"
)

const TEST_DB = "test_oolio"

func TestListProducts(t *testing.T) {
	database, err := db.ConnectDB(TEST_DB)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewRepository(database)
	seed.Seed(TEST_DB)
	products, err := repo.ListProducts(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(products) == 0 {
		t.Fatal("expected products, got none")
	}
}

func TestGetProduct(t *testing.T) {
	database, err := db.ConnectDB(TEST_DB)
	if err != nil {
		t.Fatal(err)
	}
	seed.Seed(TEST_DB)
	repo := NewRepository(database)
	product, err := repo.GetProduct(context.Background(), 10)
	if err != nil {
		t.Fatal(err)
	}
	if product.ID == "" {
		t.Fatal("expected product, got none")
	}
}

func TestPlaceOrder(t *testing.T) {
	database, err := db.ConnectDB(TEST_DB)
	if err != nil {
		t.Fatal(err)
	}
	seed.Seed(TEST_DB)
	repo := NewRepository(database)
	order := models.Order{
		Items: []models.Item{
			{
				ProductID: 10,
				Quantity:  1,
			},
		},
	}
	_, err = repo.PlaceOrder(context.Background(), order)
	if err != nil {
		t.Fatal(err)
	}
}
