package db

import (
	"database/sql"
	"log"
	//"os"
	//"path/filepath"
	_ "github.com/mattn/go-sqlite3"
)

func New() *sql.DB{
	//appPath,err := os.Executable()
	//if err != nil {
		//log.Fatal(err)
	//}
	//dbFile := filepath.Join(filepath.Dir(appPath),"scheduler.db")
	//_, err = os.Stat(dbFile)

	//var install bool
	//if err != nil {
		//install = true
	//}
	 
	db, err := sql.Open("sqlite3","scheduler.db")
	if err != nil {
		log.Fatal("init db", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("ping db", err)
	}

	return db
}