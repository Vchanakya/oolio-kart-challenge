package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Vchanakya/oolio-kart-challenge/backend/api"
	"github.com/Vchanakya/oolio-kart-challenge/backend/internal/models"
	"github.com/Vchanakya/oolio-kart-challenge/backend/internal/repository"
	"github.com/Vchanakya/oolio-kart-challenge/backend/internal/utils"
)

type Handler struct {
	server *api.Server
	repo   repository.Repository
}

func NewHandler(server *api.Server, repo repository.Repository) *Handler {
	return &Handler{
		server: server,
		repo:   repo,
	}
}

func (h *Handler) ListProducts(ctx context.Context) ([]api.Product, error) {
	products, err := h.repo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	var apiProducts []api.Product
	for _, product := range products {
		apiProducts = append(apiProducts, api.Product{
			ID:    api.NewOptString(product.ID),
			Name:  api.NewOptString(product.Name),
			Price: api.NewOptFloat32(float32(product.Price)),
			Image: api.NewOptProductImage(api.ProductImage{
				Thumbnail: api.NewOptString(product.Image.Thumbnail),
				Mobile:    api.NewOptString(product.Image.Mobile),
				Tablet:    api.NewOptString(product.Image.Tablet),
				Desktop:   api.NewOptString(product.Image.Desktop),
			}),
			Category: api.NewOptString(product.Category),
		})
	}
	return apiProducts, nil
}

func (h *Handler) GetProduct(ctx context.Context, req api.GetProductParams) (api.GetProductRes, error) {
	product, err := h.repo.GetProduct(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}
	return &api.Product{
		ID:    api.NewOptString(product.ID),
		Name:  api.NewOptString(product.Name),
		Price: api.NewOptFloat32(float32(product.Price)),
		Image: api.NewOptProductImage(api.ProductImage{
			Thumbnail: api.NewOptString(product.Image.Thumbnail),
			Mobile:    api.NewOptString(product.Image.Mobile),
			Tablet:    api.NewOptString(product.Image.Tablet),
			Desktop:   api.NewOptString(product.Image.Desktop),
		}),
		Category: api.NewOptString(product.Category),
	}, nil
}

func (h *Handler) PlaceOrder(ctx context.Context, req api.OptOrderReq) (api.PlaceOrderRes, error) {
	reqOrder, _ := req.Get()

	if reqOrder.GetCouponCode().Value != "" {
		if len(reqOrder.GetCouponCode().Value) < 8 || len(reqOrder.GetCouponCode().Value) > 10 {
			return nil, fmt.Errorf("invalid coupon code")
		}
		if !utils.HasCouponInAtLeastTwo([]string{"couponbase1.gz", "couponbase2.gz", "couponbase3.gz"}, reqOrder.GetCouponCode().Value) {
			return nil, fmt.Errorf("invalid coupon code")
		}
	}

	var items []models.Item
	for _, item := range reqOrder.GetItems() {
		productId, _ := strconv.ParseInt(item.GetProductId(), 10, 64)
		items = append(items, models.Item{
			ProductID: productId,
			Quantity:  int64(item.GetQuantity()),
		})
	}
	newOrder := models.Order{
		Items:      items,
		CouponCode: reqOrder.GetCouponCode().Value,
	}
	order, err := h.repo.PlaceOrder(ctx, newOrder)
	if err != nil {
		return nil, err
	}
	var apiItems []api.OrderItemsItem
	for _, item := range order.Items {
		apiItems = append(apiItems, api.OrderItemsItem{
			ProductId: api.NewOptString(strconv.FormatInt(item.ProductID, 10)),
			Quantity:  api.NewOptInt(int(item.Quantity)),
		})
	}
	var apiProducts []api.Product
	for _, product := range order.Products {
		apiProducts = append(apiProducts, api.Product{
			ID:    api.NewOptString(product.ID),
			Name:  api.NewOptString(product.Name),
			Price: api.NewOptFloat32(float32(product.Price)),
			Image: api.NewOptProductImage(api.ProductImage{
				Thumbnail: api.NewOptString(product.Image.Thumbnail),
				Mobile:    api.NewOptString(product.Image.Mobile),
				Tablet:    api.NewOptString(product.Image.Tablet),
				Desktop:   api.NewOptString(product.Image.Desktop),
			}),
			Category: api.NewOptString(product.Category),
		})
	}
	return &api.Order{
		ID:        api.NewOptString(order.ID),
		Total:     api.NewOptFloat64(order.Total),
		Discounts: api.NewOptFloat64(order.Discounts),
		Items:     apiItems,
		Products:  apiProducts,
	}, nil
}
