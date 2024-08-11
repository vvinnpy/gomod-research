package gomod_research

import "fmt"

func PrintVersion() {
	fmt.Println(`v2.0.0: 
	tag 修改为 v2, 但是模块版本没有修改为 github.com/forifdo/gomod-research/v2,
	**不能**通过 require github.com/forifdo/gomod-research v2.0.0 导入模块,
	**只能**通过 require github.com/forifdo/gomod-research <commit-id> 导入模块,
	go mod tidy 后, 会自动补充为 require github.com/forifdo/gomod-research v1.0.0-0.<commit-time>-<commit-id>	
	`)
}
