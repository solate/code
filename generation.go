package generation

import (
	"fmt"
	"github.com/solate/generation/config"
	"github.com/solate/generation/utils"
	"os"
	"strings"
)

////根据传入uri 生成代码
//func GenerationCodeByUri(uri string) (err error) {
//
//	//example, err := utils.ReadFile(config.Config.ServiceExample)
//	//if err != nil {
//	//	return
//	//}
//	//
//	////判断目录是否存在
//	//if !utils.IsExistFile(config.Config.ServiceExport) {
//	//	err = os.MkdirAll(config.Config.ServiceExport, 0777)
//	//	if err != nil {
//	//		return
//	//	}
//	//}
//	//
//	//err = GenerateService(example)
//	//if err != nil {
//	//	return
//	//}
//
//	return
//}

//根据模板输出代码
// template 模板字符串
// exportPath 输出路径
// fileName 文件名字, 可以是struct, 下划线等, 都会转成下划线格式的
// params 模板需要的参数
func GenerateCode(template, exportPath, fileName string, params ...interface{}) (err error) {
	templateStr := strings.TrimSpace(fmt.Sprintf(template, params...))
	err = utils.WriteFile(utils.GenerateUnderlineFileName(exportPath, fileName), templateStr)
	return
}

//根据配置config, 生成整个工程
func GenerateProjectCode() (err error) {

	//生成Service
	err = GenerateService()
	if err != nil {
		return
	}

	//生成module
	err = GenerateModule()

	return

}

func GenerateService() (err error) {

	//判断目录是否存在, 不存在就创建
	if !utils.IsExistFile(config.Config.ServiceExport) {
		err = os.MkdirAll(config.Config.ServiceExport, 0777)
		if err != nil {
			return
		}
	}

	//service 所有方法一个文件
	err = GenerateCode(config.Config.ServiceTemplate, config.Config.ServiceExport, config.Config.ServiceFileName, config.Config.ServiceTemplateParams...)

	return
}

func GenerateModule() (err error) {
	//判断目录是否存在, 不存在就创建
	if !utils.IsExistFile(config.Config.ModuleExport) {
		err = os.MkdirAll(config.Config.ModuleExport, 0777)
		if err != nil {
			return
		}
	}

	//module 每个方法一个文件
	for k, v := range config.Config.ModuleTemplateParams {
		err = GenerateCode(config.Config.ModuleTemplate, config.Config.ModuleExport, k, v...)
	}

	return
}
