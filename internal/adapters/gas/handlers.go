package gas

import (
	"fmt"
	"gas-test/internal/adapters"
	"net/http"
)

type handler struct {
	gasService Service
}

func NewHandler(service Service) adapters.Handler {
	return &handler{gasService: service}
}

func (h *handler) Statistic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}
