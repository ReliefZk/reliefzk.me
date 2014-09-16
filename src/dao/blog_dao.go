package dao

import (
	"model"
	"log"
)

type BlogDao interface {
	Save (blog *model.Blog)
	IDao
}

type BlogDaoImpl struct {
	
}

func (blogDaoImpl *BlogDaoImpl) Save(blog *model.Blog){
	log.Println(blog)
}
