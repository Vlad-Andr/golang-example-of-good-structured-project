package tests

import (
	"best-structure-example/internal/handler"
	"best-structure-example/internal/mocks"
	"best-structure-example/internal/model"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestApi_GetAllItems(t *testing.T) {
	router := chi.NewRouter()
	mockService := new(mocks.MockService)
	testHandler := handler.NewHandler(mockService, nil)

	mockService.On("ListItems").Return([]model.Item{
		{
			ID:        "123",
			Name:      "Test",
			Price:     1990,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}, nil)

	router.Get("/api/v1/items", testHandler.ListItems)

	expectedItems := []model.Item{
		{
			ID:        "123",
			Name:      "Test",
			Price:     1990,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}

	req := httptest.NewRequest("GET", "/api/v1/items", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var items []model.Item
	err := json.Unmarshal(resp.Body.Bytes(), &items)
	assert.NoError(t, err)
	assert.Equal(t, expectedItems, items)
}
