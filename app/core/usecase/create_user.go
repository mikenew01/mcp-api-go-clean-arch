package usecase

import (
	"clean-arch/core/domain"
	"clean-arch/core/repository"
	"clean-arch/infrastructure/logger"
	"context"
	"time"
)

type (
	CreateUserUseCase interface {
		Execute(context.Context, CreateUserInput) (CreateUserOutput, error)
	}

	CreateUserInput struct {
		Id    string `json:"id" validate:"required,min=1,max=32"`
		Name  string `json:"name" validate:"required,min=5,max=50"`
		Email string `json:"email" validate:"required,email"`
	}

	CreateUserPresenter interface {
		Output(domain.User) CreateUserOutput
	}

	CreateUserOutput struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	createUserUseCase struct {
		repository repository.UserRepository
		presenter  CreateUserPresenter
		ctxTimeout time.Duration
	}
)

func NewCreateUserUseCase(
	repository repository.UserRepository,
	presenter CreateUserPresenter,
	t time.Duration,
) createUserUseCase {
	return createUserUseCase{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (usecase createUserUseCase) Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	logger.StructuredLog("createUserUseCase", "Execute", "input", input).Info("Init crete user function")
	ctx, cancel := context.WithTimeout(ctx, usecase.ctxTimeout)
	defer cancel()

	userDomain := domain.NewUser(input.Id, input.Name, input.Email)

	user, err := usecase.repository.PutItem(ctx, userDomain)
	if err != nil {
		return usecase.presenter.Output(domain.User{}), err
	}
	logger.StructuredLog("createUserUseCase", "Execute", "output", user).Info("Finish crete user function")

	return usecase.presenter.Output(user), nil
}
