package controllers

import (
	"net/http"
	"GoWebMVC-CRUD/database"
	"GoWebMVC-CRUD/model"
)



//显示首页模板
func Show(w http.ResponseWriter,r*http.Request){
	t.ExecuteTemplate(w,"index",nil)
}

//查询首页数据
func Index(w http.ResponseWriter,r*http.Request){
	db := database.DbConn()
	//selDB,err := db.Query("SELECT * FORM Article ORDER BY id DESC ")
	selDB, err := db.Query("SELECT * FROM Article ORDER BY id DESC")

	if err != nil{
		panic(err.Error())
	}
	//ar := moArticle{} //没有分离前的写法
	//ra := []Article{}  //没有分离前的写法
	ar := model.Article{}   //分离后需要把包和结构体名称加进来
	ra := []model.Article{}  //分离后需要把包和结构体名称加进来
	for selDB.Next(){
		var id int
		var title ,content string
		err = selDB.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}
		ar.Id = id
		ar.Title = title
		ar.Content = content
		ra = append(ra,ar)

	}
	t.ExecuteTemplate(w,"index",ra)
	defer db.Close()
}

