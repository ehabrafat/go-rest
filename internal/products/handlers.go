package products

import (
	"myproject/internal/json"
	repo "myproject/internal/sqlc/out"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	qry *repo.Queries 
}

func NewHandler(qry *repo.Queries) *Handler {
	return &Handler {
		qry: qry,
	}
}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request){
	todos, err := h.qry.ListTodos(r.Context())
	if err != nil{
		panic(err.Error())
	}
	json.Write(w, http.StatusOK, todos);
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request){
	idStr := chi.URLParam(r, "id")

  id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }
	todo, err := h.qry.GetTodoById(r.Context(), int64(id))

	if err != nil{
		panic(err.Error())
	}
	
	json.Write(w, http.StatusOK, todo);
}