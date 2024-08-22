package data

import (
	"database/sql"
	"github.com/cloudflare/cloudflare-go"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type Service struct {
	DB  *sql.DB
	API *cloudflare.API
}

func NewDataService() (*Service, error) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}

	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		return nil, err
	}
	return &Service{
		DB:  db,
		API: api,
	}, nil
}
