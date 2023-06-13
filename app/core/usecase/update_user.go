package usecase

import (
	"clean-arch/core/domain"
	"clean-arch/core/repository"
	"clean-arch/infrastructure/logger"
	"context"
	"time"
)

type (
	UpdateUserUseCase interface {
		Execute(context.Context, UpdateUserInput) (UpdateUserOutput, error)
	}

	UpdateUserInput struct {
		Id    string `json:"id" validate:"required,min=1,max=32"`
		Name  string `json:"name" validate:"required,min=5,max=50"`
		Email string `json:"email" validate:"required,email"`
	}

	UpdateUserPresenter interface {
		Output(domain.User) UpdateUserOutput
	}

	UpdateUserOutput struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	updateUserUseCase struct {
		repository repository.UserRepository
		presenter  UpdateUserPresenter
		ctxTimeout time.Duration
	}
)

func NewUpdateUserUseCase(
	repository repository.UserRepository,
	presenter UpdateUserPresenter,
	t time.Duration,
) updateUserUseCase {
	return updateUserUseCase{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (usecase updateUserUseCase) Execute(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	logger.StructuredLog("updateUserUseCase", "Execute", "input", input).Info("Init update user function")
	ctx, cancel := context.WithTimeout(ctx, usecase.ctxTimeout)
	defer cancel()

	userDomain := domain.NewUser(input.Id, input.Name, input.Email)

	user, err := usecase.repository.UpdateItem(ctx, userDomain)
	if err != nil {
		return usecase.presenter.Output(domain.User{}), err
	}
	logger.StructuredLog("updateUserUseCase", "Execute", "output", user).Info("Finish update user function")

	return usecase.presenter.Output(user), nil
}
