package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/HorodeckiyMykhailo/go_final_project/internal/error"
)

func(h *Handler) Done(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == "" {
		error.JResponse(w,http.StatusBadRequest,"Не указан идентификатор")
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		error.JResponse(w,http.StatusBadRequest,"Некоректный формат ID")
		return 
	}

	req, err := h.repo.GetTaskById(id)
	if err != nil{
		error.JResponse(w,http.StatusBadRequest,"Задача не найдена")
		return
	}

	if req.Repeat == "" {
		rowsAffected,err := h.repo.Delete(idInt)
		if err != nil || rowsAffected == 0 {
			error.JResponse(w,http.StatusBadRequest,"Задача не найдена")
			return
		}
	} else {
		now := time.Now()
		nextDate, err := NextDate(now,req.Date,req.Repeat)
		if err != nil {
			error.JResponse(w, http.StatusBadRequest,"Неверный формат правила повторения")
		   return
		}
		rowsAffected,err := h.repo.UpdateTask(nextDate,req.Title,req.Comment,req.Repeat,idInt)
		if err != nil || rowsAffected == 0 {
			error.JResponse(w,http.StatusBadRequest,"Задача не найдена")
			return
		}
	

	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("{}"))
}