package action

import (
	"dao"
	"fmt"
	"html/template"
	"log"
	"model"
	"net/http"
	"util"
)

var userDao dao.UserDao = new(dao.UserDaoImpl)

/**
* index page handler
 */
func IndexPageHandler(writer http.ResponseWriter, reader *http.Request) {
	log.Println("index page handler.....")
	tmpl, err := util.MergeTemplate("home", nil)
	if err != nil {
		fmt.Println("html parse failed....", err)
		return
	}
	tmpl.Execute(writer, nil)
}

/**
* reg page handler
 */
func RegPageHandler(writer http.ResponseWriter, reader *http.Request) {
	log.Println("index page handler.....")
	tmpl, err := template.ParseFiles(util.Template_prefix_path + "reg.html")
	if err == nil {
		log.Println("html parse failed....")
		return
	} else {
		fmt.Println("#########", err)
	}
	tmpl.Execute(writer, nil)
}

/**
* login page handler
 */
func LoginPageHandler(writer http.ResponseWriter, reader *http.Request) {
	log.Println("index page handler.....")
	tmpl, err := template.ParseFiles(util.Template_prefix_path + "login.html")
	if err == nil {
		log.Println("html parse failed....")
		return
	}
	tmpl.Execute(writer, nil)
}

/**
* login page handler
 */
func LoginActionHandler(resp http.ResponseWriter, req *http.Request) {
	log.Println("index page handler.....")
	// fill user model
	username := req.FormValue("username")
	password := req.FormValue("password")
	var user *model.User = &model.User{Name: username, Password: password}
	userDao.Save(user)
	// template
	tmpl, err := template.ParseFiles(util.Template_prefix_path + "success.html")
	if err == nil {
		log.Println("html parse failed....")
		return
	}
	tmpl.Execute(resp, user)
}
