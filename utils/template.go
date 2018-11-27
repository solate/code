package utils

import (
	"bytes"
	"text/template"
)

//将传入的魔板，解析参数后，返回为文本
func TemplateParse(name, tmpl string, s interface{}) (str string, err error) {
	t := template.Must(template.New(name).Parse(tmpl))
	var buffer bytes.Buffer
	err = t.Execute(&buffer, s)
	return buffer.String(), err
}

//将生成的魔板放入buffer中
func BufferTemplateParse(buffer *bytes.Buffer, name, tmpl string, s interface{}) (err error) {
	t := template.Must(template.New(name).Parse(tmpl))
	return t.Execute(buffer, s)
}


