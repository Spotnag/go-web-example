package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DataService struct {
	DB *sql.DB
}

func NewDataService() (*DataService, error) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}
	return &DataService{DB: db}, nil
}
