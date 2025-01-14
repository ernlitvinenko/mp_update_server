package database

import (
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"sync"
)

var (
	instance *DatabaseApi = nil
	once                  = sync.Once{}
)

func InitializeDB() *DatabaseApi {
	once.Do(func() {
		dsn := os.Getenv("POSTGRES_DSN")
		db, err := sqlx.Connect("pgx", dsn)
		instance = &DatabaseApi{DbInstance: db}
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
