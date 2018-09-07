package utils

import "strings"

//获取文件import body 内容
//file 读取文件的全部内容
func GetImportBody(file string) (str string) {
	return GetTargetAfterPart(file, "import", "(", ")")

}

//获得结构体内容字符串
//file 读取文件全部内容
//sub 要取得结构体名称
func GetStructBody(file, sub string) (str string) {
	return GetTargetAfterPart(file, sub, "{", "}")
}

//生成下划线格式文件名,
//name 驼峰名字改下划线
//
func GenerateUnderlineFileName(path, name string) string {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return path + CamelToUnderline(name) + ".go"
}

//生成脱粉格式文件名
//name 下划线名称改脱粉
func GenerateCamelFileName(path, name string) string {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return path + UnderlineToCamel(name) + ".go"
}
