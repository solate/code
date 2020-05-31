package project

var Config *Setting

//配置文件
type Setting struct {
	//ServicePackageName string //Service 包名
	//ServiceImportBody  string //Service 导入
	//ServiceTemplate    string //Service 使用模板字符串
	//ServiceExport      string //Service 输出文件路径
	//
	//ModuleTemplate   string //Module 使用模板文件路径
	//ModuleExport     string //Module 输出文件路径
	//ModuleExample    string //Module 使用样例路径
	//ModuleName       string //Module Struct 名字
	//ModuleMethodType string //Module 方法类型: get/post/put/delete
	//ModuleUri        string //Module 请求路径  /outpatient/reports
	//ModuleNote       string //Module uri 注释

	MarkdownPath string // 文档路径地址

	//ServiceExample        string   // service 样例
	ServiceTemplate       string        // service 整体模板
	ServiceTemplateParams []interface{} // service 模板需要使用参数
	ServiceExport         string        //Service 输出文件路径
	ServiceFileName       string        //service 都写在一个文件内所以需要一个文件名

	//ModuleExample        string   // module 样例
	ModuleTemplate       string                   // module 整体层模板
	ModuleTemplateParams map[string][]interface{} // moddule 模板需要使用的参数, string 生成struct名，[]interface{} 参数
	ModuleExport         string                   //Module 输出文件路径
}

//修改配置文件
func SetConfig(cfg *Setting) {
	Config = cfg
}
