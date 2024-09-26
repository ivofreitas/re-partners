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

// CalculatePacks
// @Summary Calculate packs needed for an order
// @Description Given the total items ordered, calculate the number of packs needed
// @Tags Fulfillment
// @Accept  json
// @Produce  json
// @Param request body domain.Order true "Order information"
// @Success 200 {object} domain.Fulfillment "Number of packs required for the given order"
// @Failure 400 {object} map[string]string "Invalid input provided"
// @Router /fulfillment/items/calculate-packs [post]
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
