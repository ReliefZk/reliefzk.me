package util

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

//模版字符串
var (
	foot,
	footer,
	head,
	header string
)

//页面对应的map
var pageMap = make(map[string]string)

//模版路径前缀
var Template_prefix_path = "../src/templates/"

//初始PULL模版变量
var pullMap = make(template.FuncMap)

/**
* 初始化
 */
func init() {
	/* 初始化网页模版 */
	head, _ = ReadFile("head.html")
	header, _ = ReadFile("header.html")
	foot, _ = ReadFile("foot.html")
	footer, _ = ReadFile("footer.html")

	/* map data */
	home, _ := ReadFile("home.html")
	pageMap["home"] = home

	_404, _ := ReadFile("404.html")
	pageMap["404"] = _404

	_error, _ := ReadFile("error.html")
	pageMap["error"] = _error

	//pull 变量
	pullMap["__server_name"] = "http://localhost"
}

/**
* 1,读取html模版
* 2,合并
* 3,输出
 */
func MergeTemplate(tmplName string, funcMap template.FuncMap) (*template.Template, error) {
	//merge map
	if funcMap != nil {
		for key, value := range funcMap {
			pullMap[key] = value
		}
	}
	//模版连接
	tmpl_str := head + header + GetPageStrByName(tmplName) + footer + foot
	//parse模版
	t, _ := template.New(tmplName).Funcs(template.FuncMap(pullMap)).Parse(tmpl_str)
	return t, nil
}

/**
* 输出静态html文件
 */
func PrintStaticTempalte(writer http.ResponseWriter, tmplName string, funcMap template.FuncMap, data interface{}) {
	tmpl, err := MergeTemplate(tmplName, funcMap)
	if err != nil {
		tmpl.Execute(writer, data)
	} else {
		PrintStaticTempalte(writer, "error", nil, nil)
	}
}

/**
* 获取模版
 */
func GetPageStrByName(tmplName string) string {
	tmpls := pageMap[tmplName]
	if tmpls == "" {
		return pageMap["404"]
	}
	return tmpls
}

/**
* 读取模版
 */
func ReadFile(name string) (string, error) {
	b, err := ioutil.ReadFile(Template_prefix_path + name)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(b), nil
}
