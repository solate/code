package generation

import (
	"fmt"
	"github.com/solate/generation/utils"
	"strings"
)

func GenerationReqByFile(fileName, reqName, formart string) (str string, err error) {
	//读取文件
	fileStr, err := utils.ReadFile(fileName)
	if err != nil {
		return
	}

	list := GenerationReq(fileStr, reqName, formart)
	return strings.Join(list, "\n"), nil
}

//生成Req代码
func GenerationReq(file, sub string, format string) (list []string) {
	return GenerationReqByBody(utils.GetStructBody(file, sub), format)
}

//根据请求结构内容生成
//reqBody 请求内容字符串
func GenerationReqByBody(reqBody string, format string) (list []string) {
	//为空直接返回
	if strings.TrimSpace(reqBody) == "" {
		return
	}

	//按行读取字符串，按行拆分
	params := strings.Split(strings.TrimSpace(reqBody), "\n")

	list = make([]string, 0, len(params))
	for _, param := range params {
		words := strings.Fields(strings.TrimSpace(param)) //按单词拆分
		formatStr := fmt.Sprintf(format, utils.CamelToUnderline(words[0]), "req."+words[0])
		list = append(list, formatStr)
	}

	return
}
