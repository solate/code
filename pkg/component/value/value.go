package value

import (
	"os"
	"path/filepath"
	"strings"
)

// 项目名称，默认创建项目文件夹也必须为这个
const project = "code"

// 默认为单元测试使用的值，main需要覆盖这些值
var (
	path   = strings.Split(Pwd(), project)[0]   // 获取当前运行文件的绝对路径，然后定位到项目目录的父目录
	Root   = filepath.Join(path, project)       // 根路径
	Export = filepath.Join(path, "code/export") // 导出项目路径，可以为任意目录
)

// 获得当前运行这个文件的绝对路径
func Pwd() string {
	str, _ := os.Getwd()
	return str
}
