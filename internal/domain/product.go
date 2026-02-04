package domain

import "context"

type Product struct {
	ID          int
	Title       string
	Price       int
	Description string
	Image       string
}

type ProductRepository interface {
	FetchById(ctx context.Context, produtID int) (Product, error)
}

type ProductUsecase interface {
	GetProduct(ctx context.Context, produtID int) (Product, error)
}
