package repository

import (

	"github.com/HorodeckiyMykhailo/go_final_project/internal/task"
)

func (r *Repository) GetTasks() ([]task.TaskRequest, error) {
	rows, err := r.db.Query("SELECT ID, Date, Title, Comment, Repeat FROM scheduler ORDER BY date ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.TaskRequest
	for rows.Next() {
		var req task.TaskRequest
		if err := rows.Scan(&req.ID, &req.Date, &req.Title, &req.Comment, &req.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, req)
	}

	// Проверяем на ошибки после итерации
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}