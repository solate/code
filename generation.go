package generation

import (
	"github.com/solate/generation/config"
	"github.com/solate/generation/utils"
	"os"
)

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
