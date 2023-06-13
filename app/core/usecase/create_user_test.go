package usecase

import (
	"clean-arch/core/domain"
	"clean-arch/infrastructure/logger"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type (
	UserRepositoryMock struct {
		result domain.User
		err    error
	}

	CreateUserPresenterMock struct{}

	testTable struct {
		name     string
		input    input
		expected expected
		assert   func(t *testing.T, repositoryMock UserRepositoryMock, createUserPresenterMock CreateUserPresenterMock, input input, output CreateUserOutput, err error) bool
	}

	input struct {
		input CreateUserInput
	}

	expected struct {
		createUserOutput CreateUserOutput
		err              error
	}
)

func (mock *UserRepositoryMock) PutItem(ctx context.Context, user domain.User) (domain.User, error) {
	return mock.result, mock.err
}

func (mock *UserRepositoryMock) FindById(context.Context, string) (domain.User, error) {
	return domain.User{}, nil
}
func (mock *UserRepositoryMock) UpdateItem(context.Context, domain.User) (domain.User, error) {
	return domain.User{}, nil
}
func (mock *UserRepositoryMock) DeleteItem(context.Context, string) error {
	return nil
}

func (mock *CreateUserPresenterMock) Output(user domain.User) CreateUserOutput {
	return CreateUserOutput{Id: user.Id, Name: user.Name, Email: user.Email}
}

func TestExecute(t *testing.T) {
	logger.NewZapLogger()
	testTable := []testTable{
		{
			name: "success execution",
			input: input{
				input: CreateUserInput{
					Id:    "1",
					Name:  "Mock",
					Email: "mock@mail.com",
				},
			},
			expected: expected{
				createUserOutput: CreateUserOutput{
					Id:    "1",
					Name:  "Mock",
					Email: "mock@mail.com",
				},
				err: nil,
			},
			assert: func(t *testing.T, repositoryMock UserRepositoryMock, createUserPresenterMock CreateUserPresenterMock, input input, output CreateUserOutput, err error) bool {
				return assert.Equal(t, input.input.Id, output.Id) &&
					assert.Equal(t, input.input.Name, output.Name) &&
					assert.Equal(t, input.input.Email, output.Email)
			},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			repositoryMock := &UserRepositoryMock{result: domain.User(tt.input.input), err: tt.expected.err}
			createUserPresenterMock := &CreateUserPresenterMock{}

			service := NewCreateUserUseCase(repositoryMock, createUserPresenterMock, time.Second*10)
			output, err := service.Execute(context.TODO(), tt.input.input)

			assert.True(t, tt.assert(t, *repositoryMock, *createUserPresenterMock, tt.input, output, err))
		})
	}
}
