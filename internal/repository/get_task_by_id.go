package repository

import (
	"database/sql"

	"github.com/HorodeckiyMykhailo/go_final_project/internal/task"
)

func (r *Repository) GetTaskById(id string) (task.TaskRequest, error) {
	var req task.TaskRequest
	row := r.db.QueryRow(`SELECT ID, Date, Title, Comment, Repeat FROM scheduler WHERE ID = ?`, id)
	err := row.Scan(&req.ID, &req.Date, &req.Title, &req.Comment, &req.Repeat)
	if err != nil || err == sql.ErrNoRows {
		return req, err
	}
	return req, nil
}
