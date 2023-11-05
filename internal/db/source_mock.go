// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package db is a generated GoMock package.
package db

import (
	context "context"
	entity "music-backend-test/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserSource is a mock of UserSource interface.
type MockUserSource struct {
	ctrl     *gomock.Controller
	recorder *MockUserSourceMockRecorder
}

// MockUserSourceMockRecorder is the mock recorder for MockUserSource.
type MockUserSourceMockRecorder struct {
	mock *MockUserSource
}

// NewMockUserSource creates a new mock instance.
func NewMockUserSource(ctrl *gomock.Controller) *MockUserSource {
	mock := &MockUserSource{ctrl: ctrl}
	mock.recorder = &MockUserSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserSource) EXPECT() *MockUserSourceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserSource) CreateUser(ctx context.Context, user *entity.UserCreate) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserSourceMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserSource)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockUserSource) DeleteUser(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserSourceMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserSource)(nil).DeleteUser), ctx, id)
}

// DislikeTrack mocks base method.
func (m *MockUserSource) DislikeTrack(ctx context.Context, userId, trackId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DislikeTrack", ctx, userId, trackId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DislikeTrack indicates an expected call of DislikeTrack.
func (mr *MockUserSourceMockRecorder) DislikeTrack(ctx, userId, trackId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DislikeTrack", reflect.TypeOf((*MockUserSource)(nil).DislikeTrack), ctx, userId, trackId)
}

// GetUserById mocks base method.
func (m *MockUserSource) GetUserById(ctx context.Context, id uuid.UUID) (*entity.UserDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(*entity.UserDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserSourceMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserSource)(nil).GetUserById), ctx, id)
}

// GetUserByUsername mocks base method.
func (m *MockUserSource) GetUserByUsername(ctx context.Context, email string) (*entity.UserDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, email)
	ret0, _ := ret[0].(*entity.UserDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockUserSourceMockRecorder) GetUserByUsername(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockUserSource)(nil).GetUserByUsername), ctx, email)
}

// LikeTrack mocks base method.
func (m *MockUserSource) LikeTrack(ctx context.Context, userId, trackId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LikeTrack", ctx, userId, trackId)
	ret0, _ := ret[0].(error)
	return ret0
}

// LikeTrack indicates an expected call of LikeTrack.
func (mr *MockUserSourceMockRecorder) LikeTrack(ctx, userId, trackId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LikeTrack", reflect.TypeOf((*MockUserSource)(nil).LikeTrack), ctx, userId, trackId)
}

// ShowLikedTracks mocks base method.
func (m *MockUserSource) ShowLikedTracks(ctx context.Context, id uuid.UUID) ([]*entity.MusicDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowLikedTracks", ctx, id)
	ret0, _ := ret[0].([]*entity.MusicDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowLikedTracks indicates an expected call of ShowLikedTracks.
func (mr *MockUserSourceMockRecorder) ShowLikedTracks(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowLikedTracks", reflect.TypeOf((*MockUserSource)(nil).ShowLikedTracks), ctx, id)
}

// UpdateUser mocks base method.
func (m *MockUserSource) UpdateUser(ctx context.Context, userDB *entity.UserDB, user *entity.UserCreate) (*entity.UserDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, userDB, user)
	ret0, _ := ret[0].(*entity.UserDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserSourceMockRecorder) UpdateUser(ctx, userDB, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserSource)(nil).UpdateUser), ctx, userDB, user)
}

// MockMusicSource is a mock of MusicSource interface.
type MockMusicSource struct {
	ctrl     *gomock.Controller
	recorder *MockMusicSourceMockRecorder
}

// MockMusicSourceMockRecorder is the mock recorder for MockMusicSource.
type MockMusicSourceMockRecorder struct {
	mock *MockMusicSource
}

// NewMockMusicSource creates a new mock instance.
func NewMockMusicSource(ctrl *gomock.Controller) *MockMusicSource {
	mock := &MockMusicSource{ctrl: ctrl}
	mock.recorder = &MockMusicSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMusicSource) EXPECT() *MockMusicSourceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMusicSource) Create(ctx context.Context, musicDb *entity.MusicDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, musicDb)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMusicSourceMockRecorder) Create(ctx, musicDb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMusicSource)(nil).Create), ctx, musicDb)
}

// Delete mocks base method.
func (m *MockMusicSource) Delete(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMusicSourceMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMusicSource)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockMusicSource) Get(ctx context.Context, musicId uuid.UUID) (*entity.MusicDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, musicId)
	ret0, _ := ret[0].(*entity.MusicDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMusicSourceMockRecorder) Get(ctx, musicId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMusicSource)(nil).Get), ctx, musicId)
}

// GetAll mocks base method.
func (m *MockMusicSource) GetAll(ctx context.Context) ([]*entity.MusicDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*entity.MusicDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockMusicSourceMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockMusicSource)(nil).GetAll), ctx)
}

// GetAllSortByTime mocks base method.
func (m *MockMusicSource) GetAllSortByTime(ctx context.Context) ([]*entity.MusicDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSortByTime", ctx)
	ret0, _ := ret[0].([]*entity.MusicDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSortByTime indicates an expected call of GetAllSortByTime.
func (mr *MockMusicSourceMockRecorder) GetAllSortByTime(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSortByTime", reflect.TypeOf((*MockMusicSource)(nil).GetAllSortByTime), ctx)
}

// GetAndSortByPopular mocks base method.
func (m *MockMusicSource) GetAndSortByPopular(ctx context.Context) ([]*entity.MusicDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAndSortByPopular", ctx)
	ret0, _ := ret[0].([]*entity.MusicDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAndSortByPopular indicates an expected call of GetAndSortByPopular.
func (mr *MockMusicSourceMockRecorder) GetAndSortByPopular(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAndSortByPopular", reflect.TypeOf((*MockMusicSource)(nil).GetAndSortByPopular), ctx)
}

// Update mocks base method.
func (m *MockMusicSource) Update(ctx context.Context, musicDb *entity.MusicDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, musicDb)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMusicSourceMockRecorder) Update(ctx, musicDb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMusicSource)(nil).Update), ctx, musicDb)
}
