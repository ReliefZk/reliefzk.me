package model

import (
	"gopkg.in/mgo.v2/bson"
)

/**
* 博客实体
 */
type Blog struct {
	Id         bson.ObjectId "_id" //ID
	Category   string        //分类
	Title      string        //标题
	Logo       string        //图片地址
	ReadNum    int32         //阅读量
	CommentNum int32         //评论量
	Content    string        //博客内容
	Summary    string        //博客摘要
}
