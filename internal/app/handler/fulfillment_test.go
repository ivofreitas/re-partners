package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"re-partners/internal/domain"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
	mockCalculate func(itemsOrdered int) *domain.Fulfillment
}

func (m *mockService) Calculate(itemsOrdered int) *domain.Fulfillment {
	return m.mockCalculate(itemsOrdered)
}

func TestCalculatePacks(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name         string
		order        *domain.Order
		mockResponse *domain.Fulfillment
		mockErr      error
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name:         "Valid order request",
			order:        &domain.Order{TotalItems: 501},
			mockResponse: &domain.Fulfillment{Packs: map[int]int{500: 1, 250: 1}},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{"packs": map[string]interface{}{"500": float64(1), "250": float64(1)}},
		},
		{
			name:         "Invalid JSON in request body",
			order:        nil, // Invalid JSON
			mockResponse: nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockService{
				mockCalculate: func(itemsOrdered int) *domain.Fulfillment {
					return tt.mockResponse
				},
			}

			h := &fulfillment{
				Service: mock,
			}

			var requestBody []byte
			if tt.order != nil {
				requestBody, _ = json.Marshal(tt.order)
			} else {
				requestBody = []byte(`invalid json`)
			}
			req := httptest.NewRequest(http.MethodPost, "/fulfillment/items/calculate-packs", bytes.NewReader(requestBody))

			w := httptest.NewRecorder()

			h.CalculatePacks(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			if tt.expectedBody != nil {
				var body map[string]interface{}
				err := json.NewDecoder(resp.Body).Decode(&body)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, body)
			}
		})
	}
}
