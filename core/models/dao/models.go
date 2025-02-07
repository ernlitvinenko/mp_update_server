package dao

import "github.com/google/uuid"

type Profile struct {
	Id       uuid.UUID
	Username string
	Password string
}

type Application struct {
	Id          string    `json:"id"`
	AppName     string    `json:"appName"`
	Description *string   `json:"description"`
	Versions    []Version `json:"versions"`
}

type Version struct {
	Id          string  `json:"id" db:"id"`
	AppId       string  `db:"app_id" json:"-"`
	VersionCode int     `json:"versionCode" db:"version_code"`
	Description *string `json:"description" db:"version_description"`
	Link        string  `json:"link" db:"link"`
}

type ListApplicationDao struct {
	Id                 string  `db:"id"`
	AppName            string  `db:"name"`
	Description        *string `db:"description"`
	VersionId          string  `db:"version_id"`
	VersionCode        int     `db:"version_code"`
	VersionDescription *string `db:"version_description"`
	Link               string  `db:"version_link"`
}
