package seed

import (
	"context"

	"github.com/Vchanakya/oolio-kart-challenge/backend/db"
	"go.mongodb.org/mongo-driver/bson"
)

func Seed(dbName string) error {
	database, err := db.ConnectDB(dbName)
	if err != nil {
		return err
	}
	collection := database.Collection("products")
	if _, err := collection.DeleteMany(context.Background(), bson.D{}); err != nil {
		return err
	}
	products := []interface{}{
		bson.M{
			"product_id": 10,
			"name":       "Chicken Waffle",
			"price":      13.3,
			"category":   "Waffle",
			"image": bson.M{
				"thumbnail": "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
				"mobile":    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
				"tablet":    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
				"desktop":   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
			},
		},
	}
	if _, err := collection.InsertMany(context.Background(), products); err != nil {
		return err
	}
	return nil
}
