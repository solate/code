package utils

import "strings"

//下划线变为驼峰
func UnderlineToCamel(str string)string{
	strPortions := strings.Split(str, "_")
	for i := range strPortions{
		strB := []byte(strPortions[i])
		if strB[0] >= 'a' && strB[0] <= 'z'{
			strB[0] = strB[0] - byte(32)
			strPortions[i] = string(strB)
		}
	}
	return strings.Join(strPortions, "")
}

//驼峰变为下划线
func CamelToUnderline(str string)string{
	strB := []byte{}
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			if i > 0 {
				strB = append(strB, '_')
			}
			strB = append(strB, str[i] + byte(32))
		} else {
			strB = append(strB, str[i])
		}
	}
	return string(strB)
}