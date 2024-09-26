package app

import (
	"github.com/swaggo/http-swagger"
	"net/http"
	_ "re-partners/docs"
	"re-partners/internal/app/handler"
	"strings"
)

type Router struct {
	handler.Fulfillment
}

func NewRouter(handler handler.Fulfillment) *Router {
	return &Router{handler}
}

const (
	calculatePacks = `/fulfillment/items/calculate-packs`
)

// ServeHTTP redirect request based on http verb and url
func (t *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && strings.EqualFold(calculatePacks, r.URL.Path):
		t.CalculatePacks(w, r)
		return
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/swagger/"):
		httpSwagger.WrapHandler(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("url not found"))
		return
	}
}
