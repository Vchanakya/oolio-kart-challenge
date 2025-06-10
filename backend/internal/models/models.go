package models

type Image struct {
	Thumbnail string `bson:"thumbnail"`
	Mobile    string `bson:"mobile"`
	Tablet    string `bson:"tablet"`
	Desktop   string `bson:"desktop"`
}

type Product struct {
	ID        string  `bson:"_id"`
	ProductID int64   `bson:"product_id"`
	Name      string  `bson:"name"`
	Price     float64 `bson:"price"`
	Image     Image   `bson:"image"`
	Category  string  `bson:"category"`
}

type Item struct {
	ProductID int64 `bson:"product_id"`
	Quantity  int64 `bson:"quantity"`
}

type Order struct {
	ID         string    `bson:"_id"`
	Items      []Item    `bson:"items"`
	Products   []Product `bson:"products"`
	CouponCode string    `bson:"coupon_code"`
	Total      float64   `bson:"total"`
	Discounts  float64   `bson:"discounts"`
}
