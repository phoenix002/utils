package main

//错误的方式1
//import "goworkspace/src/tuils" 绝对路径 //main.go:3:8: package goworkspace/src/tuils is not in GOROOT (D:\Go\go\go1.16\src\goworkspace\src\tuils)

//错误的方式2
//import "./utils"  //相对路径  //main.go:6:8: cannot find package "." in:
//D:\goworkspace\src\utils\utils

//错误的方式3
//import "GO/go/go1.16/src/goworkspace/src/utils/utils"
//main.go:9:8: package GO/go/go1.16/src/goworkspace/src/utils/utils is not in GOROOT (D:\Go\go\go1.16\src\GO\go\go1.16\src\goworkspace\src\utils\utils)

import (
	"github.com/phoenix002/utils"
)

func main() {  //main.go:9:2: undefined: utils
	utils.Count()
}
