package value

import (
	"os"
	"path"
)

var (
	goPath = os.Getenv("GOPATH")
	Root   = path.Join(goPath, "/src/github.com/solate/code/")       // 根路径
	Export = path.Join(goPath, "/src/github.com/solate/code/export") //导出项目路径
)

// 获得当前运行这个文件的绝对路径
func Pwd() string {
	str, _ := os.Getwd()
	return str
}
