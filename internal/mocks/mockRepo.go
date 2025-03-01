package mocks

import (
	"best-structure-example/internal/model"
	"github.com/stretchr/testify/mock"
)

type ItemMockRepository struct {
	mock.Mock
}

func (m *ItemMockRepository) GetItem(id string) (*model.Item, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Item), args.Error(1)
}

func (m *ItemMockRepository) CreateItem(user *model.Item) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *ItemMockRepository) ListItems() ([]model.Item, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Item), args.Error(1)
}
