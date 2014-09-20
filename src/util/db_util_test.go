package util

import (
	"testing"
	"util"
)

func Test_initSession(t *testing.T) {
	session := util.InitSession()
	if session == nil {
		t.Error("初始化session失败!")
	} else {
		t.Log("session: ", session)
	}
}

//声明全局变量
var dbUtil util.DbUtil

func Test_GetSession(t *testing.T) {
	session := dbUtil.GetSession()
	if session == nil {
		t.Error("获取session失败!")
	} else {
		t.Log("session: ", session)
	}
}

func Test_ReturnResource(t *testing.T) {
	session := dbUtil.GetSession()
	if session == nil {
		t.Error("获取session失败!")
	} else {
		t.Log("session: ", session)
	}

	dbUtil.ReturnResource(session)
}

func Test_MapToBsonMParse(t *testing.T) {
	_map := make(map[string]string)
	_map["a"] = "1"
	_map["b"] = "2"
	_map["3"] = "3"

	bsonM := dbUtil.MapToBsonMParse(_map)
	t.Log("-----", bsonM)
}
