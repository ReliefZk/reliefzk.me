package util

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/**
* 初始化

func init() {
	log.Println("init............")
}
*/
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
	log.Println(t)
	if err != nil {
		log.Println(err.Error())
		return
	} else {
		t.Execute(writer, data)
	}
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
* 包含模板
 */
func ParseHtmlTemplateToStr(tname string, data interface{}) (string, error) {
	s, err := LoadHtmlTemplateToStr(tname)
	if err != nil {
		log.Println("解析模板[", tname, "]错误,", err.Error())
		return "", err
	}
	log.Println(s)
	t, err := template.New(tname).Parse(s)
	if err != nil {
		log.Println("解析模板[", tname, "]错误,", err.Error())
		return "", err
	}
	buff := bytes.NewBufferString("")
	t.Execute(buff, data)
	t_str := buff.String()
	log.Println(t_str)
	return t_str, nil
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
	//parse模版
	t := template.New(tname)
	t = t.Funcs(template.FuncMap{"template": ParseHtmlTemplateToStr})
	t, err = t.Parse(tmpl_str)
	if err != nil {
		log.Println("编译模板失败！", err.Error())
		return nil, err
	}
	t.Execute(os.Stdout, map[string]string{"Content": "bloglist.html"})
	return t, nil
}
