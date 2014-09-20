package action

import (
	"dao"
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
	util.PrintStaticTempalte(writer, "index.html", map[string]string{"Content": "common/bloglist.html"})
}

/**
* reg page handler
 */
func RegPageHandler(writer http.ResponseWriter, reader *http.Request) {
	log.Println("reg page handler.....")
	util.PrintStaticTempalte(writer, "reg", nil)
}

/**
* login page handler
 */
func LoginPageHandler(writer http.ResponseWriter, reader *http.Request) {
	log.Println("login page handler.....")
	util.PrintStaticTempalte(writer, "login", nil)
}

/**
* login action page handler
 */
func LoginActionHandler(writer http.ResponseWriter, req *http.Request) {
	log.Println("login action page handler.....")
	// fill user model
	username := req.FormValue("username")
	password := req.FormValue("password")
	var user *model.User = &model.User{Name: username, Password: password}
	userDao.Save(user)

	// template
	util.PrintStaticTempalte(writer, "success", user)
}

/**
* about page handler
 */
func CRHandler(writer http.ResponseWriter, req *http.Request) {
	log.Println("about page handler.....")
	util.PrintStaticTempalte(writer, "common/copyright.html", nil)
}
