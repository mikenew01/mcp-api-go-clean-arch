package presenter

import (
	"clean-arch/core/domain"
	"clean-arch/core/usecase"
)

type createUserPresenter struct{}

var _ usecase.CreateUserPresenter = (*createUserPresenter)(nil)

func NewCreateUserPresenter() usecase.CreateUserPresenter {
	return createUserPresenter{}
}

func (presenter createUserPresenter) Output(user domain.User) usecase.CreateUserOutput {
	return usecase.CreateUserOutput{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}
