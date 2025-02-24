package repository

import "database/sql"

func(r *Repository) Count() *sql.Row {
	query := `SELECT COUNT(*) FROM scheduler`
	row := r.db.QueryRow(query)
	return row
}