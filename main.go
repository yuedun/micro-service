package main

import (
	"fmt"
	"github.com/russross/blackfriday"
)

func main() {
	//先做一个markdown转HTML的练习和go module依赖管理
	mdstr := "# 标题一"
	input := []byte(mdstr)
	output := blackfriday.Run(input)
	html := string(output)
	fmt.Println(">>>>", html)
}
