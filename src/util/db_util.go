/**
* 数据库Util
 */
package util

import (
	"gopkg.in/mgo.v2"
	"log"
)

//数据库连接个数
const (
	MinCore = 5
)

//mongo数据库IP
var mongo_server_ip = "127.0.0.1"

//mgo.Session类型的channel
var sessionChannel = make(chan *mgo.Session, MinCore)

/**
* 初始化函数
 */
/*func init() {
	for i := 0; i < MinCore; i++ {
		sessionChannel <- InitSession()
	}
}*/

/**
* 创建session
 */
func InitSession() *mgo.Session {
	session, err := mgo.Dial(mongo_server_ip)
	if err != nil {
		log.Println("创建session失败！", err.Error())
		return nil
	}
	log.Println("create session 成功!")
	return session
}

/**
* 数据库连接帮助类
 */
type DbUtil struct {
}

/**
* 获取session
* 这可能超时
 */
func (dbUtil *DbUtil) GetSession() *mgo.Session {
	sess, ok := <-sessionChannel
	if ok {
		return sess
	}
	return nil
}

/**
* 返回session
* 这可能超时
 */
func (dbUtil *DbUtil) ReturnResource(sess *mgo.Session) {
	sessionChannel <- sess
}
