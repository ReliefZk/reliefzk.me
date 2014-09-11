package main

import (
	"util"
	"os"
)

func main() {
	t, _ := util.MergeTemplate("home.html", nil)
	t.Execute(os.Stdout, nil)
}
