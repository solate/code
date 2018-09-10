package utils

import "strings"

func GetPackageName(file string) (str string) {
	return GetTargetAfterPart(file, "package", " ", "\n")
}

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

// url类型, 生成方法名,
// 例如: /register/test => RegisterTest
// path 方法路径
// method 方法类型，get/post
// 方法名 method + path 首字母都大写
func GetUriMethodName(path string, method string) (methodName string) {

	methodName = strings.Title(strings.ToLower(method))
	words := strings.Split(path, "/") //路径以 / 分割
	for _, v := range words {
		methodName += strings.Title(v)
	}
	return
}
