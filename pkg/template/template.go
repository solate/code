package template

import (
	"github.com/solate/util/go/gofile"
	"github.com/solate/util/go/gotemplate"
	"gitlab.dove.im/im/tl-codegen/utils"
	"os"
	"path"
	"strings"
)

type Template struct {
	Name         string      // 模板名称
	TemplatePath string      // 模板文件路径
	ExportPath   string      // 输出全文件， 路径/文件名
	Structure    interface{} // 对象指针
}

// 读取模板，并将内容解析后写入到指定路径中
// p 模板对应的结构体
func (s *Template) Write() (err error) {
	template, err := gofile.ReadFile(s.TemplatePath)
	if err != nil {
		return
	}

	data, err := gotemplate.TemplateParse(s.Name, template, s.Structure)
	if err != nil {
		return
	}

	data = strings.TrimSpace(data)
	if err = gofile.WriteFile(s.ExportPath, data); err != nil {
		return
	}
	return
}

//判断目录是否存在, 不存在就创建
func (s *Template) CheckDir() (err error) {
	dir := path.Dir(s.ExportPath)
	if !utils.IsExistFile(dir) {
		err = os.MkdirAll(dir, 0777)
		return
	}
	return
}
