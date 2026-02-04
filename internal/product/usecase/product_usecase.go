package usecase

import (
	"AlurBayar/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type productUsecaseImpl struct {
	productRepository domain.ProductRepository
	redisClient       *redis.Client
}

func NewProductUsecase(repo domain.ProductRepository, redis *redis.Client) domain.ProductUsecase {
	return &productUsecaseImpl{
		productRepository: repo,
		redisClient:       redis,
	}
}

func (u *productUsecaseImpl) GetProduct(ctx context.Context, produtID int) (domain.Product, error) {
	cacheKey := fmt.Sprintf("product:%d", produtID)
	val, err := u.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var product domain.Product
		json.Unmarshal([]byte(val), &product)
		return product, nil
	}
	product, err := u.productRepository.FetchById(ctx, produtID)
	if err != nil {
		return domain.Product{}, err
	}
	productJson, _ := json.Marshal(product)
	u.redisClient.Set(ctx, cacheKey, productJson, 24*time.Hour)

	return product, nil
}
