package adapters

import "net/http"

type Handler interface {
	Statistic(w http.ResponseWriter, r *http.Request)
}
