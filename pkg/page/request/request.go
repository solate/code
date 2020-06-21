package request

import (
	"fmt"
	"strings"

	"github.com/solate/code/utils"
	"github.com/solate/util/go/gostring"
)

// 服务参数
type Request struct {
}

// 初始化
func New() *Request {
	return &Request{}
}

// 启动服务
func (s *Request) Start() (err error) {
	//TODO 启动服务代码

	return
}

// 生成
func (s *Request) Req(file, sub string, format string) (list []string) {
	format := `data.Add("%s", %s)`
	return
}

//生成Req代码
func GenerationReq(file, sub string) (list []string) {
	format := `data.Add("%s", %s)`
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
		formatStr := fmt.Sprintf(format, gostring.CamelToUnderline(words[0]), "req."+words[0])
		list = append(list, formatStr)
	}

	return
}
