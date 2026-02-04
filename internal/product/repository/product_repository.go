package repository

import (
	"AlurBayar/internal/domain"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ProductRepositoryImpl struct {
	DB     *sql.DB
	Client *http.Client
}

func NewProductRepository(DB *sql.DB) domain.ProductRepository {
	return &ProductRepositoryImpl{
		DB: DB,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (p *ProductRepositoryImpl) FetchById(ctx context.Context, produtID int) (domain.Product, error) {
	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", produtID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return domain.Product{}, err
	}
	res, err := p.Client.Do(req)
	if err != nil {
		return domain.Product{}, err
	}
	defer res.Body.Close()

	var dto FakeStoreProduct
	if err := json.NewDecoder(res.Body).Decode(&dto); err != nil {
		return domain.Product{}, err
	}

	return domain.Product{
		ID:          dto.ID,
		Title:       dto.Title,
		Price:       int(dto.Price),
		Description: dto.Description,
		Image:       dto.Image,
	}, nil
}
