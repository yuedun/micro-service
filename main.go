package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
)

func main() {
	//先做一个markdown转HTML的练习和go module依赖管理
	mdstr := "# 标题一"
	input := []byte(mdstr)
	unsafe := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Println(">>>>", string(html))
}
