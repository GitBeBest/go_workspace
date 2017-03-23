package entities

import (
	"database/sql"
)

type model interface{
	Add() sql.Result
	GetById(id int) sql.Row
	GetByAttribute(attr []string) sql.Row
	GetAllByAttribute(attr []string) sql.Rows
	DeleteById(id int) sql.Result
	DeleteByAttr(attr []string) sql.Result
	UpdateByAttr(attr []string) sql.Result

	CheckErr(err error)
}
