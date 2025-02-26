package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/HorodeckiyMykhailo/go_final_project/internal/error"
	"github.com/HorodeckiyMykhailo/go_final_project/internal/task"
)


func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	var req task.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Ошибка десериализации JSON", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		error.JResponse(w,http.StatusBadRequest,"Заголовок задачи обязателен")
		return
	}

	now := time.Now().Truncate(24 * time.Hour)
	nowFormatted := now.Format("20060102")

	if req.Date == "" {
		req.Date = nowFormatted
	} else {
		parsedDate, err := time.Parse("20060102", req.Date)
		if err != nil {
			error.JResponse(w, http.StatusBadRequest,"Некорректный формат даты")
			return
		}

		parsedDate = parsedDate.Truncate(24 * time.Hour)

		if parsedDate.Before(now) {
			if req.Repeat == "" {
				req.Date = nowFormatted
			} else {
				nextDate, err := NextDate(now, req.Date, req.Repeat)
				if err != nil {
					error.JResponse(w,http.StatusBadRequest ,"Некорректное правило повторения")
					return
				}
				req.Date = nextDate
			}
		} else {
			req.Date = parsedDate.Format("20060102")
		}
	}

	id, err := h.repo.AddTask(req.Date, req.Title, req.Comment, req.Repeat)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при добавлении задачи: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(map[string]string{"id": fmt.Sprintf("%d", id)})
}
