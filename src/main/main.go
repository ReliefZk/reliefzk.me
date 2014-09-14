package main

/**
* util.
*
 */
import (
	"action"
	"log"
	"net/http"
	"util"
)

func main() {
	/*log.Println("server start...")
	Route()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("ListenAndServe failed!", err.Error())
		return
	}*/

	str, _ := util.ParseHtmlTemplateToStr("index.html", nil)
	log.Println(str)
}

/**
* 路由信息
 */
func Route() {
	http.HandleFunc("/", action.IndexPageHandler)
	http.HandleFunc("/login", action.LoginPageHandler)
	http.HandleFunc("/login.action", action.LoginActionHandler)
	http.HandleFunc("/reg", action.RegPageHandler)
	http.HandleFunc("/reg.action", action.LoginActionHandler)
	http.HandleFunc("/about", action.AboutActionHandler)
}
