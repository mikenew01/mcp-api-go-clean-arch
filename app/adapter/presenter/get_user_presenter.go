package presenter

import (
	"clean-arch/core/domain"
	"clean-arch/core/usecase"
)

type getUserPresenter struct{}

var _ usecase.GetUserPresenter = (*getUserPresenter)(nil)

func NewGetUserPresenter() usecase.GetUserPresenter {
	return getUserPresenter{}
}

func (presenter getUserPresenter) Output(user domain.User) usecase.GetUserOutput {
	return usecase.GetUserOutput{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}
