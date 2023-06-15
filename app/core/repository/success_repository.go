package repository

import (
	"clean-arch/core/domain"
	"context"
)

type TransactionRepository interface {
	PutItem(context.Context, domain.Transaction) (domain.Transaction, error)
}
