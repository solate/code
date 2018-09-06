package generation

import (
	"fmt"
	"github.com/solate/generation/utils"
	"strings"
)

//根据给定接口字符串列表，生成代码
// interfaceStr 接口字符串
// templatePath 模板路径
// writePath 输出路径
// fileName 输出文件名, 直接使用接口名，会自动转为下划线格式
// params 模板需要使用的参数
func GenerateInterfaceCode(templatePath string, writePath, fileName string, params ...interface{}) (err error) {

	//模板添加对应参数
	writeStr, err := GenerateTemplate(templatePath, params...)
	if err != nil {
		return
	}

	err = WriteFile(writePath, fileName, writeStr)
	return
}

//拆分接口方法名和注释
func SpliteInterface(str string) (methodName, note string) {
	words := strings.Split(str, "//")   //获取注释
	end := strings.Index(words[0], "(") //获取接口名

	//根据接口获取文件名和注释
	methodName = strings.TrimSpace(words[0][:end])
	note = words[1]
	return
}

//生成模板输出
func GenerateTemplate(templatePath string, params ...interface{}) (str string, err error) {
	//读取模板
	templateStr, err := utils.ReadFile(templatePath)
	if err != nil {
		return
	}
	//模板添加对应参数
	return fmt.Sprintf(templateStr, params...), nil
}

//输出数据到文件
func WriteFile(writePath, fileName, data string) (err error) {
	//输出文件, 文件名使用下划线, 数据清除前后空格
	//在命令行中运行，不要直接在goland运行，路径不对
	err = utils.WriteFile(utils.GenerateUnderlineFileName(writePath, fileName), strings.TrimSpace(data))
	return
}
