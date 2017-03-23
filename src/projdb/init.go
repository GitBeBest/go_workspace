package projdb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"bytes"
)

var MyDb *sql.DB

func Init() *sql.DB{
	var db_connect bytes.Buffer
	db_connect.WriteString(db_user)
	db_connect.WriteString(":")
	db_connect.WriteString(db_password)
	db_connect.WriteString("@tcp(")
	db_connect.WriteString(db_host)
	db_connect.WriteString(":")
	db_connect.WriteString(db_port)
	db_connect.WriteString(")/")
	db_connect.WriteString(db_name)
	db_connect.WriteString("?charset=")
	db_connect.WriteString(db_charset)

	MyDb,_ = sql.Open(db_driver, db_connect.String())
	MyDb.SetMaxOpenConns(500)
	MyDb.SetMaxIdleConns(300)
	MyDb.Ping()

	return MyDb
}
