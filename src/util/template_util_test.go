package util

import (
	"fmt"
	"os"
	"testing"
)

func Test_ReadFile(t *testing.T) {
	s, err := ReadFile("foot.html")
	if err != nil {
		t.Error("读取模版失败!", err)
	} else {
		t.Log("模版内容：", s)
	}
}

func Test_MergeTemplate(t *testing.T) {
	tmpl, err := MergeTemplate("foot.html")
	if err != nil {
		t.Error("合并模版失败!", err)
	} else {
		t.Log("模版内容：")
		tmpl.Execute(os.Stdout, nil)
	}
}
