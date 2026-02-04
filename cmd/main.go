package main

import (
	"AlurBayar/internal/config"
	"AlurBayar/internal/product/delivery/http"
	"AlurBayar/internal/product/delivery/http/route"
	"AlurBayar/internal/product/repository"
	"AlurBayar/internal/product/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	config.LoadConfig()
	db := config.NewDB()
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: config.GetEnv("REDIS_ADDR", "localhost:6379"),
	})
	pRepo := repository.NewProductRepository(db)
	pUsecase := usecase.NewProductUsecase(pRepo, rdb)
	pHandler := http.NewProductHandler(pUsecase)

	r := gin.Default()
	route.MapProductRoutes(r, pHandler)

	// 7. Jalankan Server
	port := config.GetEnv("APP_PORT", "8080")
	log.Printf("Aplikasi AlurBayar berjalan di port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
