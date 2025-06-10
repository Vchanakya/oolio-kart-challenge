# Oolio Kart Challenge - Backend Service

A Go-based backend service for handling e-commerce operations, including product management and order processing.

## Features

- Product Management:

  - List all products
  - Get product details by ID
  - Product seeding with initial data

- Order Processing:
  - Place orders with multiple items
  - Coupon code validation (15% discount)
  - Concurrent coupon validation across multiple files
  - Order total calculation with discounts

## Tech Stack

- Go 1.22+
- MongoDB
- OpenAPI/Swagger for API documentation

## Project Structure

```
backend/
├── cmd/              # Main application entry point
│   └── main.go       # Application startup code
├── coupons/          # Coupons files
├── db/               # Database connection and configuration
│   └── db.go         # MongoDB connection
├── internal/         # Internal packages
│   ├── handler/      # API handlers
│   ├── models/       # Data models
│   ├── repository/   # Data access layer
│   ├── seed/         # Database seeding
│   └── utils/        # Utility functions
└── api/             # API definitions
```

## Getting Started

### Prerequisites

- Go 1.22+
- MongoDB 4.0+
- Make sure MongoDB is running on localhost:27017

### Installation

1. Clone the repository
2. Install dependencies:

```bash
go mod download
```

### Running the Application

```bash
go run cmd/main.go
```

The server will start on port 8080.

### API Documentation

The API is documented using OpenAPI/Swagger. You can find the API specification in `api/openapi.yaml`.

### API Endpoints

- `GET /products` - List all products
- `GET /products/{id}` - Get product details
- `POST /orders` - Place a new order

### Security

The API is protected with API key authentication. All requests must include:

```
X-API-Key: apitest
```

## Testing

The project includes unit tests for the repository layer. To run tests:

```bash
go test ./...
```

## Data Models

### Product

```go
type Product struct {
    ID        string  `bson:"_id"`
    ProductID int64   `bson:"product_id"`
    Name      string  `bson:"name"`
    Price     float64 `bson:"price"`
    Image     Image   `bson:"image"`
    Category  string  `bson:"category"`
}
```

### Order

```go
type Order struct {
    ID         string    `bson:"_id"`
    Items      []Item    `bson:"items"`
    Products   []Product `bson:"products"`
    CouponCode string    `bson:"coupon_code"`
    Total      float64   `bson:"total"`
    Discounts  float64   `bson:"discounts"`
}
```
