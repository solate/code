package {{.Package}}

import(
{{- range $key, $value := .ImportList}}
  "{{$value}}"
{{- end}}
)

// 服务参数
type {{.StructName}} struct {

}

// 初始化
func New() *{{.StructName}}  {
	return &{{.StructName}}{}
}

// 启动服务
func (s *{{.StructName}}) Start() (err error) {
	//TODO 启动服务代码

	return
}





