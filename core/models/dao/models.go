package dao

import "github.com/google/uuid"

type Profile struct {
	Id       uuid.UUID
	Username string
	Password string
}

type Application struct {
	Id       string    `json:"id"`
	AppName  string    `json:"appName"`
	Versions []Version `json:"versions"`
}

type Version struct {
	Id          string  `json:"id" db:"id"`
	AppId       string  `db:"app_id" json:"-"`
	VersionCode int     `json:"versionCode" db:"version_code"`
	Description *string `json:"description" db:"description"`
	Link        string  `json:"link" db:"link"`
}

type ListApplicationDao struct {
	Id          string `db:"id"`
	AppName     string `db:"name"`
	VersionId   string `db:"version_id"`
	VersionCode int    `db:"version_code"`
	Link        string `db:"version_link"`
}
