package products

import (
	"myproject/internal/json"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request){
	products := struct {
		Products []string 
	} {
		Products: []string {"GG"},
	}
	json.Write(w, http.StatusOK, products);
}