package lib

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"io"
	"entities"
	"strconv"
)

func Index(w http.ResponseWriter, r * http.Request)  {
	fmt.Fprintf(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(entities.GetTodo());err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "todoshow:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request)  {
	var todo entities.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil{
		panic(err)
	}

	if err := r.Body.Close();err != nil{
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil{
			panic(err)
		}
	}

	t := entities.RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(t);err != nil{
		panic(err)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	var field = new(struct{
		name string
		password string
	})
	field.name = "wangwu"
	field.password = "6543210"

	var attr = new(struct{
		id int
	})

	attr.id = 1
	entities.Update(field, attr)
	fmt.Fprintln(w, "更新成功！")
}

func GetUser(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err == nil{
		entities.GetById(id)
	}
	fmt.Fprintln(w, "获取成功！")
}
