package usecase

import (
	"context"
	"music-backend-test/internal/entity"

	"github.com/google/uuid"
)

//go:generate mockgen -source=./interfaces.go -destination=./usecases_mock.go -package=usecase
type UserInteractor interface {
	Create(ctx context.Context, user *entity.UserCreate) (uuid.UUID, error)
	GetById(ctx context.Context, id uuid.UUID) (*entity.UserDB, error)
	GetByUsername(ctx context.Context, username string) (*entity.UserDB, error)
	Update(ctx context.Context, id uuid.UUID, user *entity.UserCreate) (*entity.UserDB, error)
	Delete(ctx context.Context, id uuid.UUID) error
	LikeTrack(ctx context.Context, userId uuid.UUID, trackId uuid.UUID) error
	DislikeTrack(ctx context.Context, userId uuid.UUID, trackId uuid.UUID) error
	ShowLikedTracks(ctx context.Context, id uuid.UUID) ([]*entity.MusicDB, error)
}

type MusicInteractor interface {
	GetAll(ctx context.Context) ([]*entity.MusicDB, error)
	Get(ctx context.Context, musicId uuid.UUID) (*entity.MusicDB, error)
	GetAndSortByPopular(ctx context.Context) ([]*entity.MusicDB, error)
	GetAllSortByTime(ctx context.Context) ([]*entity.MusicDB, error)
	Create(ctx context.Context, musicCreate *entity.MusicParse) error
	Update(ctx context.Context, id uuid.UUID, musicUpdate *entity.MusicParse) error
	Delete(ctx context.Context, id uuid.UUID) error
}
