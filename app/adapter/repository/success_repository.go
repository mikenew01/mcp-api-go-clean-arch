package repository

import (
	"clean-arch/core/domain"
	"clean-arch/core/repository"
	"context"
)

type TransactionRepository struct {
	db NoSQL
}

var _ repository.TransactionRepository = (*TransactionRepository)(nil)

func NewSuccessRepositoryDynamoDB(db NoSQL) TransactionRepository {
	return TransactionRepository{
		db: db,
	}
}

func (repository TransactionRepository) PutItem(ctx context.Context, item domain.Transaction) (domain.Transaction, error) {
	successTransaction, err := repository.PutItem(ctx, item)

	if err != nil {
		return domain.Transaction{}, err
	}

	return successTransaction, nil
}
