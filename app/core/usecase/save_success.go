package usecase

import (
	"clean-arch/core/domain"
	"clean-arch/core/repository"
	"context"
	"time"
)

type (
	SaveSuccessUseCase interface {
		Execute(context.Context, SaveTransactionInput) (SaveTransactionOutput, error)
	}
	SaveTransactionInput = domain.Transaction

	SaveSuccessPresenter interface {
		Output(domain.Transaction) SaveTransactionOutput
	}
	SaveTransactionOutput = domain.Transaction

	saveTransactionUseCase struct {
		repository repository.TransactionRepository
		presenter  SaveSuccessPresenter
		ctxTimeout time.Duration
	}
)

func NewSaveSuccessUseCase(
	repository repository.TransactionRepository,
	presenter SaveSuccessPresenter,
	t time.Duration,
) saveTransactionUseCase {
	return saveTransactionUseCase{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (usecase saveTransactionUseCase) Execute(ctx context.Context, input SaveTransactionInput) (SaveTransactionOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ctxTimeout)
	defer cancel()

	transaction, err := usecase.repository.PutItem(ctx, input)
	if err != nil {
		return SaveTransactionOutput{}, err
	}

	return transaction, nil
}
