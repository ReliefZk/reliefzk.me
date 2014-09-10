package dao

import (
	"model"
	"log"
)

type UserDao interface {
	Save(user *model.User)
	IDao
}

type UserDaoImpl struct {
	
}

func (userDao *UserDaoImpl) Save(user *model.User) {
	log.Println(user)
}
