package controllers

import (
	"net/http"
	"GoWebMVC-CRUD/database"
)




func Register(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "Register", nil)
}

//注册页面
func Insert(w http.ResponseWriter, r *http.Request) {

	db := database.DbConn()   //通过导入的方法连接数据数据库
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		pwd := r.FormValue("pwd")
		insForm, err := db.Prepare("INSERT INTO Name( name,city,pwd) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, pwd)
	}
	http.Redirect(w, r, "/register", 301)
}



