package usecase

import (
	"AlurBayar/internal/domain"
	"context"
	"fmt"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type TransactionUsecaseImpl struct {
	productRepo     domain.ProductRepository
	transactionRepo domain.TransactionRepository
	serverKey       string
}

func NewTransactionUsecase(pr domain.ProductRepository, tr domain.TransactionRepository, key string) domain.TransactionUsecase {
	return &TransactionUsecaseImpl{
		productRepo:     pr,
		transactionRepo: tr,
		serverKey:       key,
	}
}

func (u *TransactionUsecaseImpl) CreateOrder(ctx context.Context, productID int) (domain.Transaction, error) {
	product, err := u.productRepo.FetchById(ctx, productID)
	if err != nil {
		return domain.Transaction{}, fmt.Errorf("produk tidak ditemukan: %v", err)
	}
	s := snap.Client{}
	s.New(u.serverKey, midtrans.Sandbox)
	orderID := fmt.Sprintf("ALUR-%d", time.Now().Unix())

	productName := product.Title
	if len(productName) > 50 {
		productName = productName[:50]
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(product.Price),
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    fmt.Sprintf("%d", product.ID),
				Name:  productName,
				Price: int64(product.Price),
				Qty:   1,
			},
		},
	}
	snapResp, errSnap := s.CreateTransaction(req)
	if errSnap != nil {
		return domain.Transaction{}, fmt.Errorf("gagal ke midtrans: %v", errSnap.GetMessage())
	}
	txData := domain.Transaction{
		ID:          orderID,
		ProductID:   product.ID,
		ProductName: productName,
		Amount:      int(product.Price),
		Status:      domain.StatusPending,
		SnapToken:   snapResp.Token,
		CreatedAt:   time.Now(),
	}
	if err := u.transactionRepo.Create(ctx, txData); err != nil {
		return domain.Transaction{}, err
	}

	return txData, nil
}

func (u *TransactionUsecaseImpl) ProcessNotification(ctx context.Context, payload map[string]interface{}) error {
	orderID, ok := payload["order_id"].(string)
	if !ok {
		return fmt.Errorf("invalid notification payload: order_id not found")
	}
	transactionStatus, ok := payload["transaction_status"].(string)
	if !ok {
		return fmt.Errorf("invalid notification payload: transaction_status not found")
	}
	var newStatus string
	switch transactionStatus {
	case "capture", "settlement":
		newStatus = domain.StatusSettlement
	case "deny", "cancel", "expire":
		newStatus = domain.StatusCancel
	case "pending":
		newStatus = domain.StatusPending
	default:
		newStatus = domain.StatusPending
	}

	err := u.transactionRepo.UpdateStatus(ctx, orderID, newStatus)
	if err != nil {
		return fmt.Errorf("failed to update transaction status: %v", err)
	}

	fmt.Printf("Transaction %s updated to %s\n", orderID, newStatus)
	return nil
}
