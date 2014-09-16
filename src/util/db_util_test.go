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

var dbUtil = new(util.DbUtil)

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
