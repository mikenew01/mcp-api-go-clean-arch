package repository

import "context"

type NoSQL interface {
	FindById(ctx context.Context, id string) (UserEntity, error)
	PutItem(ctx context.Context, item UserEntity) (UserEntity, error)
	UpdateItem(ctx context.Context, item UserEntity) (UserEntity, error)
	DeleteItem(ctx context.Context, id string) error
}
