package handler

import (
	"net/http"
	"strconv"

	"github.com/HorodeckiyMykhailo/go_final_project/internal/error"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == "" {
		error.JResponse(w,"Не указан идентификатор")
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		error.JResponse(w,"Некоректный формат ID")
		return 
	}

	rowsAffected,err := h.repo.Delete(idInt)
	if err != nil || rowsAffected == 0 {
		error.JResponse(w,"Задача не найдена")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("{}"))
}