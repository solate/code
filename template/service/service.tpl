package services

import (
	"github.com/solate/code/pkg/component/logger"
)

//创建{{.StructName}}CreatHandle
func {{.StructName}}CreatHandle({{.StructName}} *models.{{.StructName}}) (err error) {
	logger.Logger.Debug("In {{.StructName}}CreatHandle")

	//因为id相同，更新或新增
	if {{.StructName}}.NewRecord() { //更新
		err = {{.StructName}}.Save()
	} else {
		err = {{.StructName}}.Create()
	}
	return
}

//删除{{.StructName}}, 设置状态为删除
func {{.StructName}}DeleteHandle(username string) (err error) {
	logger.Logger.Debug("In {{.StructName}}DeleteHandle")

	param := map[string]interface{}{
		"is_delete": models.Delete,
	}
	{{.StructName}} := new(models.{{.StructName}})
	err = {{.StructName}}.UpdateByUsername(username, param)
	return
}

//获得列表
func {{.StructName}}ListHandle(name, roleNames string, page, perPage int) (list []models.{{.StructName}}, count int, err error) {
	logger.Logger.Debug("In {{.StructName}}ListHandle")

	{{.StructName}} := models.{{.StructName}}{}
	list, count, err = {{.StructName}}.FindList(name, roleNames, page, perPage)
	return
}
