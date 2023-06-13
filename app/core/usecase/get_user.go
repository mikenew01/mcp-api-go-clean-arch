package usecase

import (
	"clean-arch/core/domain"
	"clean-arch/core/repository"
	"clean-arch/infrastructure/logger"
	"context"
	"time"
)

type (
	GetUserUseCase interface {
		Execute(context.Context, GetUserInput) (GetUserOutput, error)
	}

	GetUserInput struct {
		Id string `json:"id"`
	}

	GetUserPresenter interface {
		Output(domain.User) GetUserOutput
	}

	GetUserOutput struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	getUserUseCase struct {
		repository repository.UserRepository
		presenter  GetUserPresenter
		ctxTimeout time.Duration
	}
)

func NewGetUserUseCase(
	repository repository.UserRepository,
	presenter GetUserPresenter,
	t time.Duration,
) getUserUseCase {
	return getUserUseCase{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (usecase getUserUseCase) Execute(ctx context.Context, input GetUserInput) (GetUserOutput, error) {
	logger.StructuredLog("getUserUseCase", "Execute", "input", input).Info("Init get user function")
	ctx, cancel := context.WithTimeout(ctx, usecase.ctxTimeout)
	defer cancel()

	user, err := usecase.repository.FindById(ctx, input.Id)
	if err != nil {
		return usecase.presenter.Output(domain.User{}), err
	}
	if user.Id == "" {
		return GetUserOutput{}, domain.ErrUserNotFound
	}
	logger.StructuredLog("getUserUseCase", "Execute", "output", user).Info("Finish get user function")

	return usecase.presenter.Output(user), nil
}
