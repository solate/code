package template

import (
	"github.com/google/uuid"
	"github.com/solate/code/pkg/component/value"
	"github.com/solate/util/go/gofile"
	"github.com/solate/util/go/gotemplate"
	"gitlab.dove.im/im/tl-codegen/utils"
	"os"
	"path"
	"strings"
)

type Template struct {
	name   string      // 模板名称，默认为structName
	Path   string      // 模板文件路径
	Export string      // 输出全文件， 路径/文件名
	Option interface{} // 对象指针, 模板需要的参数
}

func New(templatePath, exportPath string, o *Option) *Template {
	return &Template{
		name:   uuid.New().String(), //唯一创建一个模板
		Path:   path.Join(value.Root, templatePath),
		Export: path.Join(value.Export, exportPath),
		Option: o,
	}
}

// 启动
func (s *Template) Start() (err error) {
	if err = s.CheckDir(); err != nil { //判断文件夹是否创建
		return
	}
	return s.Write()
}

// 读取模板，并将内容解析后写入到指定路径中
// p 模板对应的结构体
func (s *Template) Write() (err error) {
	template, err := gofile.ReadFile(s.Path)
	if err != nil {
		return
	}

	data, err := gotemplate.TemplateParse(s.name, template, s.Option)
	if err != nil {
		return
	}

	data = strings.TrimSpace(data)
	if err = gofile.WriteFile(s.Export, data); err != nil {
		return
	}
	return
}

//判断目录是否存在, 不存在就创建
func (s *Template) CheckDir() (err error) {
	dir := path.Dir(s.Export)
	if !utils.IsExistFile(dir) {
		err = os.MkdirAll(dir, 0777)
		return
	}
	return
}

//// 驼峰改为下划线格式的文件名称
//func (s *Template) CamelToUnderlineFile(path, fileName string) string {
//	return filepath.Join(path, gostring.CamelToUnderline(fileName), suffix)
//}
