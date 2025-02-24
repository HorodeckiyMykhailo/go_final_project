package handler

import (
	"encoding/json"
	"net/http"
	"github.com/HorodeckiyMykhailo/go_final_project/internal/error"
)

func(h *Handler) GetTasks(w http.ResponseWriter, r *http.Request){
	var count int
	err := h.repo.Count().Scan(&count)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if count == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"tasks": []interface{}{}})
		return
	}

	tasks, err := h.repo.GetTasks()
	if err != nil {
		error.JResponse(w,"Ошибка при получении задач")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"tasks": tasks})
}