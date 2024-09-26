package handler

import (
	"encoding/json"
	"net/http"
	"re-partners/internal/app/service"
	"re-partners/internal/domain"
)

type Fulfillment interface {
	CalculatePacks(w http.ResponseWriter, r *http.Request)
}

type fulfillment struct {
	Service service.Fulfillment
}

type Marshal func(v any) ([]byte, error)

func New(service service.Fulfillment) Fulfillment {
	return &fulfillment{service}
}

func (f *fulfillment) CalculatePacks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)

	var (
		err    error
		order  *domain.Order
		result *domain.Fulfillment
	)
	defer func() {
		if err != nil {
			encoder.Encode(map[string]string{"error": err.Error()})
			return
		}
		encoder.Encode(result)
	}()

	if err = json.NewDecoder(r.Body).Decode(&order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result = f.Service.Calculate(order.TotalItems)

	w.WriteHeader(http.StatusOK)
}
