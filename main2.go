package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func main() {
	//先做一个markdown转HTML的练习和go module依赖管理
	mdstr := "# 标题一"
	input := []byte(mdstr)
	unsafe := blackfriday.Run(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	fmt.Println(">>>>", string(html))
}
