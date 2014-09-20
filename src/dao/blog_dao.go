package dao

import (
	"log"
	"model"
	"util"
)

//博客数据库
var BLOG_DB = "my_blog"

//数据库管理帮助类
var dbUtil = new(DbUtil)

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
	msession := dbUtil.GetSession()
	defer dbUtil.ReturnResource(msession)
	cursor := msession.DB(BLOG_DB).C("blog")

	bson.M bsonM = dbUtil.MapToBsonMParse(queryParams)

	// 读取数据
    blogList := []model.Blog{}
    err := cursor.Find(&bsonM).All(&blogList)
    if err != nil {
    	log.Println("mongo query error.", err.Error())
    }

    return blogList
}
