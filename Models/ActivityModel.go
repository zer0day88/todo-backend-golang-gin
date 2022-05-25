package Models

import (
	"todo/backend/Utils"
)

type Activity struct {
	Id         uint            `json:"id" gorm:"index;primaryKey; autoIncrement"`
	Email      string          `json:"email"`
	Title      string          `json:"title"`
	Created_at Utils.JSONTime  `json:"created_at" gorm:"autoCreateTime:milli"`
	Updated_at Utils.JSONTime  `json:"updated_at" gorm:"autoUpdateTime:milli"`
	Deleted_at *Utils.JSONTime `json:"deleted_at"`
}

type Activity_model struct {
	Id         uint    `json:"id"`
	Email      string  `json:"email"`
	Title      string  `json:"title"`
	Created_at string  `json:"created_at"`
	Updated_at string  `json:"updated_at"`
	Deleted_at *string `json:"deleted_at"`
}

type Activity_Created struct {
	Created_at Utils.JSONTime `json:"created_at"`
	Updated_at Utils.JSONTime `json:"updated_at"`
	Id         uint           `json:"id"`
	Title      string         `json:"title"`
	Email      string         `json:"email"`
}
