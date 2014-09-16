package util

import (
	"testing"
)

func Test_MergeTemplate(t *testing.T) {
	tmpl, err := ParseTemplate("index.html")
	if err != nil {
		t.Error("合并模版失败!", err)
	} else {
		t.Log("模版内容：", tmpl)
		//tmpl.Execute(os.Stdout, map[string]string{"Content": "bloglist.html"})
	}
}

func Test_ParseHtmlTemplateToStr(t *testing.T) {
	str := MyTemplateDirective("index.html", nil)
	t.Log("str = ", str)
}
