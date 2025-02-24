package repository

func(r *Repository) Delete(id int)(int64,error){
	query := `DELETE FROM scheduler WHERE id = ?`
	res, err := r.db.Exec(query,id)
	if err != nil {
		return 0,err
	}
	return res.RowsAffected()
}