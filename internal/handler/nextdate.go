package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	
	

	_ "github.com/mattn/go-sqlite3"
)



func NextDate(now time.Time, date string, repeat string) (string, error) {
	//Проверка на наличие правила повторения
	if repeat == "" {
		return "", errors.New("отсутствует правило повторения")
	}

	//Проверка формата даты
	taskDate, err := time.Parse("20060102", date)
	if err != nil {
		return "", fmt.Errorf("некорректный формат даты: %w", err)
	}

	//Правило повторения при "d 1"
	if repeat == "d 1" && !taskDate.After(now) {
		return now.Format("20060102"), nil
	}

	//проверка соответствия условиям правила повторения
	rpt := strings.Split(repeat, " ")
	for {
		if rpt[0] == "y" && len(rpt) < 2 {
			taskDate = taskDate.AddDate(1, 0, 0)
		} else if rpt[0] == "d" && len(rpt) == 2 {
			days, err := strconv.Atoi(rpt[1])
			if err != nil || days < 1 || days > 400 {
				return "", fmt.Errorf("некорректное число дней: d %s", rpt[1])
			}
			taskDate = taskDate.AddDate(0, 0, days)
		} else {
			return "", fmt.Errorf("неподдерживаемый формат правила повторения: %s", repeat)
		}

		if taskDate.After(now) {
			return taskDate.Format("20060102"), nil
		}
	}

}

func(h *Handler) NextDateHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры запроса
	nowStr := r.URL.Query().Get("now")
	dateStr := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	if nowStr == "" || dateStr == "" {
		http.Error(w, "missing required parameters", http.StatusBadRequest)
		return
	}

	// Парсим текущее время
	now, err := time.Parse("20060102", nowStr)
	if err != nil {
		http.Error(w, "invalid now format", http.StatusBadRequest)
		return
	}

	// Вычисляем следующую дату
	nextDate, err := NextDate(now, dateStr, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Отправляем ответ
	fmt.Fprint(w, nextDate)
}












