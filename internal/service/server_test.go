package service

import (
	"best-structure-example/internal/mocks"
	"best-structure-example/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestServer_GetItems(t *testing.T) {
	mockRepo := new(mocks.ItemMockRepository)
	itemService := NewService(mockRepo, nil, nil)

	expectedItem1 := &model.Item{
		ID:        "123",
		Name:      "Test1",
		Price:     1890,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	expectedItem2 := &model.Item{
		ID:        "1234",
		Name:      "Test2",
		Price:     1899,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	mockRepo.On("ListItems").Return([]model.Item{*expectedItem1, *expectedItem2}, nil)

	items, err := itemService.ListItems()

	assert.NoError(t, err)
	assert.Equal(t, []model.Item{*expectedItem1, *expectedItem2}, items)
	mockRepo.AssertExpectations(t)
}

func TestServer_GetItem(t *testing.T) {
	mockRepo := new(mocks.ItemMockRepository)
	itemService := NewService(mockRepo, nil, nil)

	expectedItem := &model.Item{
		ID:        "123",
		Name:      "Test",
		Price:     1890,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	mockRepo.On("GetItem", "123").Return(expectedItem, nil)

	item, err := itemService.GetItem("123")

	assert.NoError(t, err)
	assert.Equal(t, expectedItem, item)
	mockRepo.AssertExpectations(t)
}

func TestServer_CreateItem(t *testing.T) {
	mockRepo := new(mocks.ItemMockRepository)
	mockKafkaProducer := new(mocks.MockKafkaProducer)
	itemService := NewService(mockRepo, mockKafkaProducer, nil)

	item := &model.Item{
		Name:  "Test",
		Price: 1890,
	}

	mockRepo.On("CreateItem", mock.Anything).Return(nil)
	mockKafkaProducer.On("Produce", mock.Anything, mock.Anything).Return(nil)
	resultItem, err := itemService.CreateItem(*item)

	assert.NoError(t, err)
	assert.Equal(t, item.Name, resultItem.Name)
	mockRepo.AssertExpectations(t)
}
