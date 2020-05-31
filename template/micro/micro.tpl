package {{.Package}}

import(
{{- range $key, $value := .ImportList}}
  "{{$value}}"
{{- end}}
)

// {{.Struct.Note}}
type {{.Struct.Name}} struct {
    *template.Option
}

// 初始化
func New(o *template.Option) *{{.Struct.Name}}  {
	return &{{.Struct.Name}}{o}
}

// 默认参数初始化
func NewDefault(name string) *{{.Struct.Name}} {
	o := &template.Option{
		Package:    name,
		ImportList: []string{},
		Struct: template.Struct{
			Note:      "",
			Name:      gostring.Ucfirst(name),
			FieldList: nil,
			FuncList:  nil,
		},
	}
	return &{{.Struct.Name}}{o}
}


// 启动服务 //TODO 为统一生产用下面这个代码，后面直接删除下面
func (s *{{.Struct.Name}}) Start() (err error) {
	//TODO 启动服务代码

	return
}


// 启动服务
func (s *{{.Struct.Name}}) Start() (err error) {
	logger.Logger.Debug("In Start")
	t := template.New(
		templateRelativePath,
		path.Join(exportDir, s.Option.Package, s.Option.Package+".go"),
		s.Option,
	)
	return t.Start()
}





