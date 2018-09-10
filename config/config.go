package config

var Config *Setting = &Setting{
	ServiceTemplate:   "./template/service.tpl",
	ServiceExport:     "./export/service/",
	ServiceExample:    "./source/service.go",
	ServiceName:       "Pro",
	ServiceMethodType: "GET",
	ServiceUri:        "/outpatient/reports",
	ServiceNote:       "测试注释",

	ModelTemplate: "./template/model.tpl",
	ModelExport:   "./export/model/",
	ModelExample:  "./source/model.go",
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

	ModelTemplate string //Model 使用模板文件路径
	ModelExport   string //Model 输出文件路径
	ModelExample  string //Model 使用样例路径

}

//修改配置文件
func SetConfig(cfg *Setting) {
	Config = cfg
}
