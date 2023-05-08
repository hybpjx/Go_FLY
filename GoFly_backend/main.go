package main

import (
	"gofly/cmd"
)

/*
@title Go-web 开发
@version 0.0.1
@description ts+gin 开发实战
*/
func main() {
	// 结束时清理缓存
	defer cmd.Clean()
	// 开始
	cmd.Start()
}
