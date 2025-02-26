package handler

import (
	"encoding/json"
	"net/http"

	"github.com/HorodeckiyMykhailo/go_final_project/internal/error"
)


func (h *Handler) GetTaskById(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == "" {
		error.JResponse(w,http.StatusBadRequest,"Не указан идентификатор")
		return
	}

	task, err := h.repo.GetTaskById(id)
	if err != nil{
		error.JResponse(w,http.StatusBadRequest,"Задача не найдена")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
	}

}