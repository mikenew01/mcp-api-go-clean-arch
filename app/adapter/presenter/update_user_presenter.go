package presenter

import (
	"clean-arch/core/domain"
	"clean-arch/core/usecase"
)

type updateUserPresenter struct{}

var _ usecase.UpdateUserPresenter = (*updateUserPresenter)(nil)

func NewUpdateUserPresenter() usecase.UpdateUserPresenter {
	return updateUserPresenter{}
}

func (presenter updateUserPresenter) Output(user domain.User) usecase.UpdateUserOutput {
	return usecase.UpdateUserOutput{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}
