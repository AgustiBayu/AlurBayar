package main

import (
	"AlurBayar/internal/config"
	tsHandler "AlurBayar/internal/transaction/delivery/http"
	tsRoute "AlurBayar/internal/transaction/delivery/http/route"
	tsRepo "AlurBayar/internal/transaction/repository"
	tsUsecase "AlurBayar/internal/transaction/usecase"

	productHandler "AlurBayar/internal/product/delivery/http"
	productRoute "AlurBayar/internal/product/delivery/http/route"
	productRepo "AlurBayar/internal/product/repository"
	productUsecase "AlurBayar/internal/product/usecase"
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
	pRepo := productRepo.NewProductRepository(db)
	pUsecase := productUsecase.NewProductUsecase(pRepo, rdb)
	pHandler := productHandler.NewProductHandler(pUsecase)

	midtransServerKey := config.GetEnv("MIDTRANS_SERVER_KEY", "isi-server-key-sandbox-kamu")

	txRepo := tsRepo.NewTransactionRepository(db)
	txUsecase := tsUsecase.NewTransactionUsecase(pRepo, txRepo, midtransServerKey)
	txHandler := tsHandler.NewTransactionHandler(txUsecase)

	r := gin.Default()
	productRoute.MapProductRoutes(r, pHandler)
	tsRoute.MapTransactionRoutes(r, txHandler)

	// 7. Jalankan Server
	port := config.GetEnv("APP_PORT", "8080")
	log.Printf("Aplikasi AlurBayar berjalan di port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
