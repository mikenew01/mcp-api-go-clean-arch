package usecase

import (
	"clean-arch/core/repository"
	"clean-arch/infrastructure/logger"
	"context"
	"time"
)

type (
	DeleteUserUseCase interface {
		Execute(context.Context, DeleteUserInput) (DeleteUserOutput, error)
	}

	DeleteUserInput struct {
		Id string `json:"id"`
	}

	DeleteUserPresenter interface {
		Output(string) DeleteUserOutput
	}

	DeleteUserOutput struct {
		Id string `json:"id"`
	}

	deleteUserUseCase struct {
		repository repository.UserRepository
		presenter  DeleteUserPresenter
		ctxTimeout time.Duration
	}
)

func NewDeleteUserUseCase(
	repository repository.UserRepository,
	presenter DeleteUserPresenter,
	t time.Duration,
) deleteUserUseCase {
	return deleteUserUseCase{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (usecase deleteUserUseCase) Execute(ctx context.Context, input DeleteUserInput) (DeleteUserOutput, error) {
	logger.StructuredLog("deleteUserUseCase", "Execute", "input", input).Info("Init delete user function")
	ctx, cancel := context.WithTimeout(ctx, usecase.ctxTimeout)
	defer cancel()

	err := usecase.repository.DeleteItem(ctx, input.Id)
	if err != nil {
		return usecase.presenter.Output(""), err
	}
	logger.StructuredLog("deleteUserUseCase", "Execute", "output", input.Id).Info("Finish delete user function")

	return usecase.presenter.Output(input.Id), nil
}
