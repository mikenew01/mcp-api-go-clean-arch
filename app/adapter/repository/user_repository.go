package repository

import (
	"clean-arch/core/domain"
	"clean-arch/core/repository"
	"clean-arch/infrastructure/logger"
	"context"
)

type UserEntity struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository struct {
	db NoSQL
}

var _ repository.UserRepository = (*UserRepository)(nil)

func NewUserRepositoryDynamoDB(db NoSQL) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repository UserRepository) FindById(ctx context.Context, userId string) (domain.User, error) {
	logger.Infof("S=UserRepositoryDynamoDB, M=FindById, stage=init, userId=%s", userId)
	out, err := repository.db.FindById(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{Id: out.Id, Name: out.Name, Email: out.Email}
	if err != nil {
		return domain.User{}, err
	}
	logger.Infof("S=UserRepositoryDynamoDB, M=FindById, stage=finish, user=%s", user)
	return user, nil
}

func (repository UserRepository) PutItem(ctx context.Context, item domain.User) (domain.User, error) {
	logger.StructuredLog("UserRepository", "PutItem", "input", item).Info("Init create user function")
	userEntity := UserEntity{Id: item.Id, Name: item.Name, Email: item.Email}
	out, err := repository.db.PutItem(ctx, userEntity)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.NewUser(out.Id, out.Name, out.Email)

	logger.StructuredLog("UserRepository", "PutItem", "input", item).Info("Finish update user function")
	return user, nil
}

func (repository UserRepository) UpdateItem(ctx context.Context, item domain.User) (domain.User, error) {
	logger.StructuredLog("UserRepository", "UpdateItem", "input", item).Info("Init update user function")
	userEntity := UserEntity{Id: item.Id, Name: item.Name, Email: item.Email}
	out, err := repository.db.UpdateItem(ctx, userEntity)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.NewUser(out.Id, out.Name, out.Email)

	logger.StructuredLog("UserRepository", "UpdateItem", "output", item).Info("Finish update user function")
	return user, nil
}

func (repository UserRepository) DeleteItem(ctx context.Context, id string) error {
	logger.StructuredLog("UserRepository", "DeleteItem", "input", id).Info("Init delete user function")
	err := repository.db.DeleteItem(ctx, id)
	if err != nil {
		return err
	}
	logger.StructuredLog("UserRepository", "DeleteItem", "output", id).Info("Finish delete user function")
	return nil
}
