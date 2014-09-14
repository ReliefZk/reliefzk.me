package util

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

/**
* 初始化
 */
var PullFuncMap = make(template.FuncMap)

func init() {
	log.Println("init............")
	PullFuncMap["mytemplate"] = MyTemplateDirective
	PullFuncMap["hostname"] = GetHost
	PullFuncMap["unescaped"] = unescaped
}

func unescaped(x string) interface{} {
	return template.HTML(x)
}

/**
* 返回服务器地址
 */
func GetHost() string {
	return "http://localhost"
}

/**
* 输出静态html文件
 */
func PrintStaticTempalte(writer http.ResponseWriter, tmplName string, data interface{}) {
	t, err := ParseTemplate(tmplName)
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		t.Execute(writer, data)
	}
}

/**
* 编译html为go模板
 */
func ParseTemplate(tname string) (*template.Template, error) {
	tmpl_str, err := LoadHtmlTemplateToStr(tname)
	if err != nil {
		log.Println("加载模板失败，", err.Error())
		return nil, err
	}
	//second compile html string
	//parse模版
	t, err := template.New(tname).Funcs(template.FuncMap(PullFuncMap)).Parse(tmpl_str)
	if err != nil {
		log.Println("编译模板失败！", err.Error())
		return nil, err
	}
	return t, nil
}

//模版路径前缀
var Template_prefix_path = "/home/zk/code/go/src/reliefzk.me/src/templates/"

/**
* 通过ioutil读取html模板
 */
func LoadHtmlTemplateToStr(tname string) (string, error) {
	sbyte, err := ioutil.ReadFile(Template_prefix_path + tname)
	if err != nil {
		return "", err
	}
	return string(sbyte), nil
}

/**
* 包含模板 directive
 */
func MyTemplateDirective(tname string, data interface{}) template.HTML {
	s, err := LoadHtmlTemplateToStr(tname)
	if err != nil {
		log.Println("解析模板[", tname, "]错误,", err.Error())
		return template.HTML("")
	}
	t, err := template.New(tname).Funcs(template.FuncMap(PullFuncMap)).Parse(s)
	if err != nil {
		log.Println("解析模板[", tname, "]错误,", err.Error())
		return template.HTML("")
	}
	buff := bytes.NewBufferString("")
	t.Execute(buff, data)
	return template.HTML(buff.String())
}
