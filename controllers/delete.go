package controllers

import (
	"log"
	"net/http"
	"GoWebMVC-CRUD/database"
)

//删除操作
func Delete(w http.ResponseWriter,r*http.Request){
	db := database.DbConn()
	ari :=r.URL.Query().Get("id")
	delForm,err := db.Prepare("DELETE FROM Article WHERE id = ?")

	if err !=nil{
		panic(err.Error())
	}
	delForm.Exec(ari)
	log.Println("delete 成功")
	defer db.Close()
	http.Redirect(w,r,"/index",301)
}
