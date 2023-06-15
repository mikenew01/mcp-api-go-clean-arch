package presenter

import (
	"clean-arch/core/domain"
	"clean-arch/core/usecase"
)

type saveSuccessPresenter struct{}

var _ usecase.SaveSuccessPresenter = (*saveSuccessPresenter)(nil)

func NewSaveTransactionPresenter() usecase.SaveSuccessPresenter {
	return saveSuccessPresenter{}
}

func (presenter saveSuccessPresenter) Output(transaction domain.Transaction) usecase.SaveTransactionOutput {
	// Aqui você pode formatar a saída como quiser.
	// No seu caso, como a saída é a mesma que a entrada, você pode simplesmente retornar a entrada como saída.
	return transaction
}
