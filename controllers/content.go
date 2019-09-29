package controllers

import (
	"net/http"
	"GoWebMVC-CRUD/database"
	"GoWebMVC-CRUD/model"
)

//查看详情
func Contentt(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Article WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := model.Article{}
	for selDB.Next() {
		var id int
		var title, content string
		err = selDB.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Title = title
		emp.Content = content
	}
	t.ExecuteTemplate(w, "content", emp)
	defer db.Close()
}
