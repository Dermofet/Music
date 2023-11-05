package repository

import (
	"context"
	"fmt"
	"music-backend-test/internal/db"
	"music-backend-test/internal/entity"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func Test_userRepository_Create(t *testing.T) {
	type fields struct {
		source *db.MockUserSource
	}
	type args struct {
		ctx  context.Context
		user *entity.UserCreate
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		setup   func(a args, f fields)
		wantErr bool
	}{
		{
			name: "success: Create userRepository",
			args: args{
				ctx: context.Background(),
				user: &entity.UserCreate{
					Username: "John",
					Password: "qwerty1234",
				},
			},
			want: uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
			setup: func(a args, f fields) {
				f.source.EXPECT().CreateUser(a.ctx, a.user).Return(uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"), nil)
			},
			wantErr: false,
		},
		{
			name: "error: Create userRepository",
			args: args{
				ctx: context.Background(),
				user: &entity.UserCreate{
					Username: "John",
					Password: "qwerty1234",
				},
			},
			want: uuid.Nil,
			setup: func(a args, f fields) {
				f.source.EXPECT().CreateUser(a.ctx, a.user).Return(nil, fmt.Errorf("can't create user in source"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				source: db.NewMockUserSource(ctrl),
			}

			r := NewUserRepository(f.source)

			tt.setup(tt.args, f)

			got, err := r.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_GetById(t *testing.T) {
	type fields struct {
		source *db.MockUserSource
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		setup   func(a args, f fields)
		wantErr bool
	}{
		{
			name: "success: GetById userRepository",
			args: args{
				ctx: context.Background(),
				id:  uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
			},
			want: uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
			setup: func(a args, f fields) {
				f.source.EXPECT().GetUserById(a.ctx, a.id).Return(&entity.UserDB{
					ID: uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "error: GetById userRepository",
			args: args{
				ctx: context.Background(),
				id:  uuid.Nil,
			},
			want: uuid.Nil,
			setup: func(a args, f fields) {
				f.source.EXPECT().GetUserById(a.ctx, a.id).Return(uuid.Nil, fmt.Errorf("can't get user from source"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				source: db.NewMockUserSource(ctrl),
			}

			r := NewUserRepository(f.source)

			tt.setup(tt.args, f)

			got, err := r.GetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_GetByUsername(t *testing.T) {
	type fields struct {
		source *db.MockUserSource
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		setup   func(a args, f fields)
		wantErr bool
	}{
		{
			name: "success: GetByUsername userRepository",
			args: args{
				ctx:      context.Background(),
				username: "John",
			},
			want: uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
			setup: func(a args, f fields) {
				f.source.EXPECT().GetUserByUsername(a.ctx, a.username).Return(&entity.UserDB{
					ID: uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "error: GetByUsername userRepository",
			args: args{
				ctx:      context.Background(),
				username: "",
			},
			want: uuid.Nil,
			setup: func(a args, f fields) {
				f.source.EXPECT().GetUserByUsername(a.ctx, a.username).Return(nil, fmt.Errorf("can't get user from source"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				source: db.NewMockUserSource(ctrl),
			}

			r := NewUserRepository(f.source)

			tt.setup(tt.args, f)

			got, err := r.GetByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.GetByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Update(t *testing.T) {
	type fields struct {
		source *db.MockUserSource
	}
	type args struct {
		ctx  context.Context
		id   uuid.UUID
		user *entity.UserCreate
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.UserDB
		setup   func(a args, f fields)
		wantErr bool
	}{
		{
			name: "success: Update userRepository",
			args: args{
				ctx: context.Background(),
				user: &entity.UserCreate{
					Username: "Paul",
					Password: "",
				},
			},
			want: &entity.UserDB{
				ID:       uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
				Username: "Paul",
				Role:     "USER",
			},
			setup: func(a args, f fields) {
				f.source.EXPECT().GetUserById(a.ctx, a.id).Return(&entity.UserDB{
					ID:       uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
					Username: "John",
				}, nil)
				f.source.EXPECT().UpdateUser(a.ctx, a.id, a.user).Return(&entity.UserDB{
					ID:       uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
					Username: "Paul",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "error: Update userRepository",
			args: args{
				ctx: context.Background(),
				user: &entity.UserCreate{
					Username: "Paul",
					Password: "",
				},
			},
			want: nil,
			setup: func(a args, f fields) {
				f.source.EXPECT().GetUserById(a.ctx, a.id).Return(nil, fmt.Errorf("can't get user from source"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				source: db.NewMockUserSource(ctrl),
			}

			r := NewUserRepository(f.source)

			tt.setup(tt.args, f)

			got, err := r.Update(tt.args.ctx, tt.args.id, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Delete(t *testing.T) {
	type fields struct {
		source *db.MockUserSource
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    error
		setup   func(a args, f fields)
		wantErr bool
	}{
		{
			name: "success: Delete userRepository",
			args: args{
				ctx: context.Background(),
				id:  uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
			},
			want: nil,
			setup: func(a args, f fields) {
				f.source.EXPECT().GetUserById(a.ctx, a.id).Return(&entity.UserDB{
					ID:       uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
					Username: "John",
				}, nil)
				f.source.EXPECT().DeleteUser(a.ctx, a.id).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "error: Delete userRepository",
			args: args{
				ctx: context.Background(),
				id:  uuid.MustParse("4a6e104d-9d7f-45ff-8de6-37993d709522"),
			},
			want: fmt.Errorf("can't delete user in source"),
			setup: func(a args, f fields) {
				f.source.EXPECT().GetUserById(a.ctx, a.id).Return(nil, fmt.Errorf("can't get user from source"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				source: db.NewMockUserSource(ctrl),
			}

			r := NewUserRepository(f.source)

			tt.setup(tt.args, f)

			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
