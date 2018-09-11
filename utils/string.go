package utils

import "strings"

//下划线变为驼峰
func UnderlineToCamel(str string) string {
	strPortions := strings.Split(str, "_")
	for i := range strPortions {
		strB := []byte(strPortions[i])
		if strB[0] >= 'a' && strB[0] <= 'z' {
			strB[0] = strB[0] - byte(32)
			strPortions[i] = string(strB)
		}
	}
	return strings.Join(strPortions, "")
}

//驼峰变为下划线
func CamelToUnderline(str string) string {
	strB := []byte{}
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			if i > 0 {
				strB = append(strB, '_')
			}
			strB = append(strB, str[i]+byte(32))
		} else {
			strB = append(strB, str[i])
		}
	}
	return string(strB)
}

//在全文中找到目标字符串，并返回目标后开始标记和结束标记之间的字符串
func GetTargetAfterPart(source, target, startTag, endTag string) (str string) {

	//如果输入或要取结构为空，直接返回空
	if strings.TrimSpace(source) == "" || strings.TrimSpace(target) == "" {
		return
	}

	//定位目标单子位置下标
	index := strings.Index(source, target)
	//如果未找到,返回空
	if index < 0 {
		return
	}

	//从子串将后面的切下来,然后根据子串定位
	subStr := source[index:]

	str = SubString(subStr, startTag, endTag)
	return
}

//获得切割后的字符串
func SubString(source, startTag, endTag string) (str string) {
	//定位请求参数内容位置, 前后{} 里的内容
	begin := strings.Index(source, startTag)
	end := strings.Index(source, endTag)

	//获得结构{}内包含的字符型
	str = source[begin+1 : end] //+1是为了不包含startTag
	return
}

func SplitByTag(source, tag string) (list []string) {
	source = strings.TrimSpace(source)        //清除空格
	listLen := strings.Count(source, tag) + 1 //2个文档就分成3份
	parts := strings.SplitN(source, tag, listLen)
	list = parts[1:] //将第一个没用部分抛弃
	return
}
