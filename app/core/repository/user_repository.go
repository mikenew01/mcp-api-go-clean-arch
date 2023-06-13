package repository

import (
	"clean-arch/core/domain"
	"context"
)

type UserRepository interface {
	FindById(context.Context, string) (domain.User, error)
	PutItem(context.Context, domain.User) (domain.User, error)
	UpdateItem(context.Context, domain.User) (domain.User, error)
	DeleteItem(context.Context, string) error
}
