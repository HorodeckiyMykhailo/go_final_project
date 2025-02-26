package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/HorodeckiyMykhailo/go_final_project/internal/error"
	"github.com/HorodeckiyMykhailo/go_final_project/internal/task"
)

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var req task.TaskRequest

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
		http.Error(w,"Ошибка",http.StatusBadRequest)
        return
    }

	if req.ID == "" {
		error.JResponse(w, http.StatusBadRequest,"Не указан идентификатор")
		return
	}

	idInt, err := strconv.Atoi(req.ID)
	if err != nil {
		error.JResponse(w,http.StatusBadRequest,"Некоректный формат ID")
		return 
	}

	if req.Title == "" {
		error.JResponse(w,http.StatusBadRequest,"Заголовок задачи обязателен")
		return
	}

	now := time.Now()
	 if req.Date == ""{
	 	req.Date = now.Format("20060102")
	 } else {
	 	parseDate, err := time.Parse("20060102",req.Date)
	 	if err != nil {
	 		error.JResponse(w,http.StatusBadRequest,"Неверный формат даты")
	 		return
	 	}
	 	if parseDate.Before(now) && req.Repeat != "" {
	 		nextDate, err := NextDate(now,req.Date,req.Repeat)
	 		if err != nil {
	 			error.JResponse(w,http.StatusBadRequest, "Неверный формат правила повторения")
				return
	 		}
	 		req.Date = nextDate

	 	}
	}

	rowsAffected,err := h.repo.UpdateTask(req.Date,req.Title,req.Comment,req.Repeat,idInt)
	if err != nil || rowsAffected == 0 {
		error.JResponse(w,http.StatusBadRequest,"Задача не найдена")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("{}"))


}
