package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	DB *sql.DB
}

func NewDataService() (*Service, error) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}

	return &Service{
		DB: db,
	}, nil
}
