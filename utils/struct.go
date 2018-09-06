package utils

import "strings"

//获得结构体内容字符串
//file 读取文件全部内容
//sub 要取得结构体名称
func GetStructBody(file, sub string) (str string) {

	//如果输入或要取结构为空，直接返回空
	if strings.TrimSpace(file) == "" || strings.TrimSpace(sub) == "" {
		return
	}

	//定位目标单子位置下标
	subBegin := strings.Index(file, sub)

	//从子串将后面的切下来
	subStr := file[subBegin:]

	//定位请求参数内容位置, 前后{} 里的内容
	begin := strings.Index(subStr, "{")
	end := strings.Index(subStr, "}")

	//获得结构{}内包含的字符型
	str = subStr[begin+1 : end] //+1是为了把前面包的{ 去掉

	return
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
