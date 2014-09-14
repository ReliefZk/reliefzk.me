package dao

import (
	"log"
	"model"
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
