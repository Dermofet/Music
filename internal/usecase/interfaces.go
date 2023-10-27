package usecase

import (
	"context"
	"music-backend-test/internal/entity"
)

//go:generate mockgen -source=./interfaces.go -destination=./usecases_mock.go -package=usecase
type UserInteractor interface {
	Create(ctx context.Context, user *entity.UserCreate) (*entity.UserID, error)
	GetById(ctx context.Context, id *entity.UserID) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, id *entity.UserID, user *entity.UserCreate) (*entity.User, error)
	Delete(ctx context.Context, id *entity.UserID) error
}