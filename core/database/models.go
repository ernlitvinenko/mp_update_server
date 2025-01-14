package database

import "github.com/jmoiron/sqlx"

type api interface {
}

type DatabaseApi struct {
	DbInstance *sqlx.DB
}
