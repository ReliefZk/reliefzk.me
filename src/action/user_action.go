package action

import (
	"dao"
	"html/template"
	"log"
	"model"
	"net/http"
	"fmt"
)

var userDao dao.UserDao = new(dao.UserDaoImpl)

/**
* index page handler
 */
func IndexPageHandler(writer http.ResponseWriter, reader *http.Request) {
	log.Println("index page handler.....")
	tmpl, err := template.ParseFiles("../src/templates/index.html")
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
	tmpl, err := template.ParseFiles("../src/templates/reg.html")
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
	tmpl, err := template.ParseFiles("src/templates/login.html")
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
	var user *model.User = &model.User{ Name: username, Password: password }
	userDao.Save(user)
	// template
	tmpl, err := template.ParseFiles("src/templates/success.html")
	if err == nil {
		log.Println("html parse failed....")
		return
	}
	tmpl.Execute(resp, user)
}
