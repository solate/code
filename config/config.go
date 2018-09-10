package config

var Config *Setting = &Setting{
	ServiceTemplate:   "./template/service.tpl",
	ServiceExport:     "./export/service/",
	ServiceExample:    "./source/service.go",
	ServiceName:       "Pro",
	ServiceMethodType: "GET",
	ServiceUri:        "/outpatient/reports",
	ServiceNote:       "测试注释",

	ModuleTemplate:   "./template/model.tpl",
	ModuleExport:     "./export/model/",
	ModuleExample:    "./source/model.go",
	ModuleName:       "Pro",
	ModuleMethodType: "GET",
	ModuleUri:        "/outpatient/reports",
	ModuleNote:       "测试注释",
}

//配置文件
type Setting struct {
	ServiceTemplate   string //Service 使用模板文件路径
	ServiceExport     string //Service 输出文件路径
	ServiceExample    string //Service 使用样例路径
	ServiceName       string //Service Struct 名字
	ServiceMethodType string //Service 方法类型: get/post/put/delete
	ServiceUri        string //Service 请求路径  /outpatient/reports
	ServiceNote       string //Service uri 注释

	ModuleTemplate   string //Module 使用模板文件路径
	ModuleExport     string //Module 输出文件路径
	ModuleExample    string //Module 使用样例路径
	ModuleName       string //Module Struct 名字
	ModuleMethodType string //Module 方法类型: get/post/put/delete
	ModuleUri        string //Module 请求路径  /outpatient/reports
	ModuleNote       string //Module uri 注释

}

//修改配置文件
func SetConfig(cfg *Setting) {
	Config = cfg
}
