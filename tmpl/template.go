package tmpl

import (
	"bytes"
	"html/template"
)

//将传入的魔板，解析参数后，返回为文本
func TemplateParse(name, tmpl string, s interface{}) (str string, err error) {
	t := template.Must(template.New(name).Parse(tmpl))
	var buffer bytes.Buffer
	err = t.Execute(&buffer, s)
	return buffer.String(), err
}
