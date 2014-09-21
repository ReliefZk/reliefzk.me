package main

/**
* util.
*
 */
import (
	"action"
	"log"
	"net/http"
)

func main() {
	log.Println("server start...")
	Route()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("ListenAndServe failed!", err.Error())
		return
	}
}

/**
* 路由信息
 */
func Route() {
	dir := "/home/zk/code/go/src/reliefzk.me/src/templates/static"
	http.Handle("/css/", http.FileServer(http.Dir(dir)))
	http.Handle("/js/", http.FileServer(http.Dir(dir)))
	http.Handle("/fonts/", http.FileServer(http.Dir(dir)))
	http.Handle("/images/", http.FileServer(http.Dir(dir)))

	//request route
	http.HandleFunc("/", action.IndexPageHandler)
	//backportal edit new blog
	http.HandleFunc("/new_blog", action.EditBlog)
}
