package repository

import (
	"AlurBayar/internal/domain"
	"context"
	"database/sql"
)

type TransactionRepositoryImpl struct {
	DB *sql.DB
}

func NewTransactionRepository(DB *sql.DB) domain.TransactionRepository {
	return &TransactionRepositoryImpl{
		DB: DB,
	}
}

func (r *TransactionRepositoryImpl) Create(ctx context.Context, ts domain.Transaction) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	status := domain.StatusPending
	SQL := `INSERT INTO transactions (id, product_id, product_name, amount, status, snap_token, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
	if _, err := tx.ExecContext(ctx, SQL,
		ts.ID,
		ts.ProductID,
		ts.ProductName,
		ts.Amount,
		status,
		ts.SnapToken,
		ts.CreatedAt); err != nil {
		return err
	}
	return tx.Commit()
}

func (r *TransactionRepositoryImpl) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE transactions SET status = $1 WHERE id = $2`
	_, err := r.DB.ExecContext(ctx, query, status, id)
	return err
}
func (r *TransactionRepositoryImpl) GetByID(ctx context.Context, id string) (domain.Transaction, error) {
	var ts domain.Transaction
	SQL := `SELECT id, product_id, product_name, amount, status, snap_token, created_at 
              FROM transactions WHERE id = $1`

	err := r.DB.QueryRowContext(ctx, SQL, id).Scan(
		&ts.ID,
		&ts.ProductID,
		&ts.ProductName,
		&ts.Amount,
		&ts.Status,
		&ts.SnapToken,
		&ts.CreatedAt,
	)
	if err != nil {
		return domain.Transaction{}, err
	}
	return ts, nil
}
