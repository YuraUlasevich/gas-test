package gas

import (
	"encoding/json"
	"gas-test/internal/adapters"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type handler struct {
	gasService Service
}

func NewHandler(service Service) adapters.Handler {
	return &handler{gasService: service}
}

func (h *handler) Statistic(w http.ResponseWriter, r *http.Request) {
	var gasData RequestData
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &gasData)
	if err != nil {
		log.Printf("Error unmarshal body: %v", err)
		http.Error(w, "can't unmarshal body", http.StatusBadRequest)
		return
	}
	var wg sync.WaitGroup
	var result ResponseData
	wg.Add(4)
	go h.gasService.AveragePricePerDay(&wg, &gasData, &result)
	go h.gasService.FrequencyDistribution(&wg, &gasData, &result)
	go h.gasService.SpentPerMonth(&wg, &gasData, &result)
	go h.gasService.TotalСost(&wg, &gasData, &result)
	wg.Wait()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
