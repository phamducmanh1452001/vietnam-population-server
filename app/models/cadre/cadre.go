package models

import "database/sql"

type Cadre struct {
	Code       string
	Name       sql.NullString
	Password   string
	Age        sql.NullInt16
	Phone      sql.NullString
	Email      sql.NullString
	SuperCode  string
	Permission int
}
