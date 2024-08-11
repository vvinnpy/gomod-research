package gomod_research

import "fmt"

func PrintVersion() {
	fmt.Println(`v2.3.0: root
	根目录和 /v2子目录都有 go.mod 文件,
	且声明的模块名都相同, 这是错误的 (无法使用此模块版本)
	`)
}
