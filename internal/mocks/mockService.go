package mocks

import (
	"best-structure-example/internal/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (s *MockService) CreateItem(item model.Item) (*model.Item, error) {
	args := s.Called(item)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Item), args.Error(1)
}

func (s *MockService) GetItem(id string) (*model.Item, error) {
	args := s.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Item), args.Error(1)
}

func (s *MockService) ListItems() ([]model.Item, error) {
	args := s.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Item), args.Error(1)
}

func (s *MockService) ProcessMessage(ctx context.Context, message []byte) error {
	args := s.Called(ctx, message)
	return args.Error(0)
}
