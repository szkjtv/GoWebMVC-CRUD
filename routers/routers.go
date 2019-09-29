package routers

import (
	"net/http"
	"GoWebMVC-CRUD/controllers"
)

func Router() {
	//启动静态文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/add", controllers.Article1)
	http.HandleFunc("/adr", controllers.Addarticle)
	http.HandleFunc("/index", controllers.Index)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/content", controllers.Contentt)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers. Delete)

	//端口需要放在路由最后预防端口堵塞路由信息
	http.ListenAndServe(":8007", nil)
}
