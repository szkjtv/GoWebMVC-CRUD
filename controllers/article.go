package controllers

import (
	"net/http"
	"GoWebMVC-CRUD/database"
)



//加载文章模板
func Article1(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "add", nil)
}


//添加文章
func Addarticle(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn()
	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")
		insForm, err := db.Prepare("INSERT INTO Article(title,content)VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(title, content)
		http.Redirect(w, r, "/add", 301)
	}

}

