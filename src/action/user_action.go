package action

import (
	"dao"
	"log"
//	"model"
	"net/http"
	"util"
)

var password = "zhoukui198752"

var userDao dao.UserDao = new(dao.UserDaoImpl)

/**
* index page handler
 */
func IndexPageHandler(writer http.ResponseWriter, reader *http.Request) {
	log.Println("index page handler.....")
	util.PrintStaticTempalte(writer, "index.html", map[string]string{"Content": "common/bloglist.html"})
}

/**
* 用户编辑博客
 *   需要校验密码
*/
func EditBlog(writer http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	log.Println("====", req.Method)
	//save a ready blog
	if req.Method == "POST" {

	} else { // edit new blog's template
		passwd := req.FormValue("passwd")
		if passwd == password {
			util.PrintStaticTempalte(writer, "common/new_blog.html", nil)
		}
	}
}
