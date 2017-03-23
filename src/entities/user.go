package entities


import (
	"time"
	"projdb"
	"fmt"
	"database/sql"
	"bytes"
	"reflect"
)

var table_name string = "user_test"
var soft_delete bool = true
//type User struct {
//	id int `json:"id"`
//	username string `json:"username"`
//	role int `json:"role"`
//	name string `json:"name"`
//	password string `json:"password"`
//	user_key string `json:"userkey"`
//	salt string `json:"salt"`
//	password_row string `json:"password_row"`
//	user_key_raw string `json:"user_key_raw"`
//	terms string `json:"terms"`
//	date_activated time.Time `json:"date_activated"`
//	last_login_time time.Time `json:"last_login_time"`
//	date_created time.Time `json:"date_created"`
//	date_updated time.Time `json:"date_updated"`
//	date_deleted time.Time `json:"date_deleted"`
//}

type User struct {
	id int `json:"id"`
	name string `json:"name"`
	password string `json:"password"`
	role int `json:"role"`
	date_created time.Time `json:"date_created"`
	date_updated time.Time `json:"date_updated"`
	date_deleted time.Time `json:"date_deleted"`
}

/**
 * 添加用户
 */
func (e *User) Add()  sql.Result{
	var query bytes.Buffer

	query.WriteString("INSERT ");
	query.WriteString(table_name)
	query.WriteString("(name, password, role) VALUES(?, ?, ?)")
	stmt, err := projdb.MyDb.Prepare(query.String())
	query.Reset()
	checkErr(err)
	res, err := stmt.Exec(e.name, e.password, e.role)
	checkErr(err)
	id,err := res.LastInsertId()
	checkErr(err)
	fmt.Print(id)
	return res
}

/**
 * 根据id获取用户
 */
func GetById(id int)  *sql.Row{
	var query bytes.Buffer

	query.WriteString("SELECT ")
	query.WriteString("id, name, password, role, date_created, date_updated ")
	query.WriteString("FROM ")
	query.WriteString(table_name)
	query.WriteString(" where id = ? and date_deleted IS NULL;")
	row:= projdb.MyDb.QueryRow(query.String(), id)
	query.Reset()
	return row
}

/**
 * 根据id删除用户
 */
func Delete(id int)  int64{
	var query bytes.Buffer

	if soft_delete {
		query.WriteString("UPDATE ")
		query.WriteString(table_name)
		query.WriteString(" SET date_deleted = ")
		query.WriteString(time.Now().String())
	}else {
		query.WriteString("DELETE FROM")
		query.WriteString(table_name)
	}

	query.WriteString(" WHERE id = ?")
	query.WriteString(";")

	stmt, err := projdb.MyDb.Prepare(query.String())
	checkErr(err)
	res, err := stmt.Exec(id)
	query.Reset()
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	return num
}

/**
 * 更新
 */
func  Update(field interface{}, attr interface{}) int64{
	var query bytes.Buffer


	immutable := reflect.ValueOf(field).Elem()
	option_attr := reflect.ValueOf(attr).Elem()

	var op_len int = immutable.NumField()
	var im_len int = option_attr.NumField()
	var index int = 0

	query_params_string := make([]interface{}, op_len + im_len)
	query.WriteString("UPDATE ")
	query.WriteString(table_name)
	query.WriteString(" SET ")

	for i :=0; i< op_len;i++  {
		query.WriteString(immutable.Type().Field(i).Name)
		query.WriteString("=?")
		if i < op_len -1 {
			query.WriteString(",")
		}
		if immutable.Field(i).Type().String() == "string" {
			query_params_string[index] = immutable.Field(i).String()
		}else{
			query_params_string[index] = immutable.Field(i).Int()
		}
		index++
	}


	for j := 0; j< im_len;j++  {
		query.WriteString(" WHERE ")
		query.WriteString(option_attr.Type().Field(j).Name)
		query.WriteString("=?")
		if j < im_len -1 {
			query.WriteString(" AND ")
		}

		if option_attr.Field(j).Type().String() == "string" {
			query_params_string[index] = option_attr.Field(j).String()
		}else{
			query_params_string[index] = option_attr.Field(j).Int()
		}
		index++
	}

	index = 0
	query.WriteString(";")
	stmt, err := projdb.MyDb.Prepare(query.String())
	checkErr(err)
	res, err := stmt.Exec(query_params_string...)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	stmt.Close();
	return num
}

/**
 * 检查错误
 */
func checkErr(err error)  {
	if err != nil{
		panic(err)
	}
}



