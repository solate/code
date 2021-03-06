package markdown

import (
	"strings"
	"testing"
)

func TestParseMarkdown(t *testing.T) {

	md := New("./api.md")
	list, err := md.ParseMarkdownByFile()
	if err != nil {
		t.Error(err.Error())
		return

	}
	t.Log(list)

	str := GetTemplateMethodDeclaration(list)
	t.Log(str)

}

func TestGenerateMarkdownTable(t *testing.T) {

	str := `| 名称   | 类型 | 说明 | 是否必填 |  默认值 |
|-------|------|------| ------ | ------- |
| outpatient_num   | string | 门诊号 | 是 |       |
| pay_num   | string | 缴费单据号 | 是 |       |
| pay_type   | string | 费用类型 | 是 |       |
| doctor_name   | string | 医生姓名 | 是 |       |
| fee   | string | 金额 | 是 |  精确到分     |
| item_name   | string | 项目名称 | 是 |       |`

	list := MarkdownTable(str, 1, 2, 3)
	t.Log(strings.Join(list, "\n"))
}
