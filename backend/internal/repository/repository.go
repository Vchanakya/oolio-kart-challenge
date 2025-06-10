package repository

import (
	"context"

	"github.com/Vchanakya/oolio-kart-challenge/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	ListProducts(ctx context.Context) ([]models.Product, error)
	GetProduct(ctx context.Context, id int64) (models.Product, error)
	PlaceOrder(ctx context.Context, order models.Order) (models.Order, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) ListProducts(ctx context.Context) ([]models.Product, error) {
	cursor, err := r.db.Collection("products").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var products []models.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) GetProduct(ctx context.Context, id int64) (models.Product, error) {
	var product models.Product
	if err := r.db.Collection("products").FindOne(ctx, bson.D{{Key: "product_id", Value: id}}).Decode(&product); err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (r *repository) PlaceOrder(ctx context.Context, order models.Order) (models.Order, error) {
	objectID := primitive.NewObjectID()
	order.ID = objectID.Hex()
	for _, item := range order.Items {
		product, err := r.GetProduct(ctx, item.ProductID)
		if err != nil {
			return models.Order{}, err
		}
		order.Products = append(order.Products, product)
		order.Total += product.Price * float64(item.Quantity)
	}
	order.Discounts = 0
	if order.CouponCode != "" {
		order.Discounts = 15
		order.Total -= order.Total * 0.15
	}
	if _, err := r.db.Collection("orders").InsertOne(ctx, order); err != nil {
		return models.Order{}, err
	}
	return order, nil
}
