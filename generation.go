package generation

import (
	"fmt"
	"github.com/solate/generation/config"
	"github.com/solate/generation/model"
	"github.com/solate/generation/utils"
	"os"
	"strings"
)

//生成代码, 这个方法会markdown读取后的所有代码都生成
func GenerationCodeByMarkDown(list []model.MarkDown) {
	exampleService, err := utils.ReadFile(config.Config.ServiceExample)
	if err != nil {
		return
	}

	//判断services目录是否存在
	if !utils.IsExistFile(config.Config.ServiceExport) {
		err = os.MkdirAll(config.Config.ServiceExport, 0777)
		if err != nil {
			return
		}
	}

	err = GenerateService(exampleService, list)
	if err != nil {
		return
	}

	//判断Module目录是否存在
	if !utils.IsExistFile(config.Config.ModuleExport) {
		err = os.MkdirAll(config.Config.ModuleExport, 0777)
		if err != nil {
			return
		}
	}

	exampleModule, err := utils.ReadFile(config.Config.ModuleExample)
	if err != nil {
		return
	}

	err = GenerateModule(exampleModule, list)

}

//根据传入uri 生成代码
func GenerationCodeByUri(uri string) (err error) {

	example, err := utils.ReadFile(config.Config.ServiceExample)
	if err != nil {
		return
	}

	//判断目录是否存在
	if !utils.IsExistFile(config.Config.ServiceExport) {
		err = os.MkdirAll(config.Config.ServiceExport, 0777)
		if err != nil {
			return
		}
	}

	err = GenerateService(example)
	if err != nil {
		return
	}

	return
}

//根据模板输出代码
// template 模板字符串
// exportPath 输出路径
// fileName 文件名字, 可以是struct, 下划线等, 都会转成下划线格式的
// params 模板需要的参数
func GenerationCode(template, exportPath, fileName string, params ...interface{}) (err error) {
	templateStr := strings.TrimSpace(fmt.Sprintf(template, params...))
	err = utils.WriteFile(utils.GenerateUnderlineFileName(exportPath, fileName), templateStr)
	return
}
