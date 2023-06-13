package presenter

import (
	"clean-arch/core/usecase"
)

type deleteUserPresenter struct{}

var _ usecase.DeleteUserPresenter = (*deleteUserPresenter)(nil)

func NewDeleteUserPresenter() usecase.DeleteUserPresenter {
	return deleteUserPresenter{}
}

func (presenter deleteUserPresenter) Output(userId string) usecase.DeleteUserOutput {
	return usecase.DeleteUserOutput{
		Id: userId,
	}
}
