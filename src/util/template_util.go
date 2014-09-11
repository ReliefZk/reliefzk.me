package util

import (
	"html/template"
	"io/ioutil"
	"log"
)

var (
	foot,
	footer,
	head,
	header string
)

var pageMap = make(map[string]string)

/**
* 初始化
 */
func init() {
	/* 初始化网页模版 */
	foot, _ = ReadFile("foot.html")
	footer, _ = ReadFile("footer.html")
	head, _ = ReadFile("head.html")
	header, _ = ReadFile("header.html")
	/* map data */
	home, _ := ReadFile("home.html")
	pageMap["home"] = home
}
//模版路径前缀
var template_prefix_path = "src/templates/"

/**
* 1,读取html模版
* 2,合并
* 3,输出
 */
func MergeTemplate(tmplName string, pullMap template.FuncMap) (*template.Template, error) {
	//模版连接
	tmpl_str := head + header + pageMap[tmplName] + footer + foot
	//parse模版
	t, _ := template.New(tmplName).Funcs(template.FuncMap(pullMap)).Parse(tmpl_str)
	return t, nil
}

/**
* 读取模版
 */
func ReadFile(name string) (string, error) {
	b, err := ioutil.ReadFile(template_prefix_path + name)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(b), nil
}
