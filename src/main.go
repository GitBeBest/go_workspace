package main

import (
	"projdb"
	"lib"
	"log"
	"net/http"
)

func addUser()  {
	//user := new(entities.User)
	//user.Name = "zhangsan"
	//user.Password = "123456"
	//user.Role = 1
	//user.Add()
}
func main() {
	projdb.Init()
	//
	//var id int
	//var name string
	//var password string
	//var role int
	//var date_created string
	//var date_updated string
	//var result *sql.Row = userModel.GetById(1)
	//result.Scan(&id, &name, &password, &role, &date_created, &date_updated)
	//fmt.Print(date_created)
	router := lib.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
