package dao

import (
	"log"
	"model"
	"util"
)

//博客数据库
var BLOG_DB = "my_blog"

//数据库管理帮助类
var dbUtil = new(util.DbUtil)

type BlogDao interface {
	Save(blog *model.Blog)
	QueryForList(queryParams map[string]string)
	IDao
}

type BlogDaoImpl struct {
}

func (blogDaoImpl *BlogDaoImpl) Save(blog *model.Blog) {
	msession := dbUtil.GetSession()
	defer dbUtil.ReturnResource(msession)
	c := msession.DB(BLOG_DB).C("blog")
	c.Insert(blog)
}

func (blogDaoImpl *BlogDaoImpl) QueryForList(queryParams map[string]string) (model.Blog[]) {
	
}
