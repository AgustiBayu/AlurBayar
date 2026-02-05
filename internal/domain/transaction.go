package domain

import (
	"context"
	"time"
)

const (
	StatusPending    = "PENDING"
	StatusSettlement = "SETTLEMENT"
	StatusCapture    = "CAPTURE"
	StatusExpire     = "EXPIRE"
	StatusCancel     = "CANCEL"
	StatusDeny       = "DENY"
	StatusRefund     = "REFUND"
)

type Transaction struct {
	ID          string
	ProductID   int
	ProductName string
	Amount      int
	Status      string
	SnapToken   string
	CreatedAt   time.Time
}

type TransactionRepository interface {
	Create(ctx context.Context, ts Transaction) error
	UpdateStatus(ctx context.Context, id string, status string) error
	GetByID(ctx context.Context, id string) (Transaction, error)
}

type TransactionUsecase interface {
	CreateOrder(ctx context.Context, productID int) (Transaction, error)
	ProcessNotification(ctx context.Context, payload map[string]interface{}) error
}
