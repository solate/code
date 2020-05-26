package controllers


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/solate/code/pkg/component/logger"
	"github.com/solate/util/gin/reply"
	"github.com/solate/util/project/errorUtil"
	"go_bms_web/services"
	"strings"
)

//{{.StructName}}Create 创建
func {{.StructName}}Create(c *gin.Context) {
	logger.Logger.Debug("In {{.StructName}}Create")

	var reqParams models.{{.StructName}}
	//调用services 错误
	err := services.{{.StructName}}CreatHandle(&reqParams)
	if err != nil {
		logger.Logger.Error(err.Error())
		reply.Response(c, reply.SetCodeMessage(errorUtil.ErrServer, "服务端处理报错"))
		return
	}
	reply.Response(c)
}

// {{.StructName}}Update 更新update
func {{.StructName}}Update(c *gin.Context) {

	reply.Response(c)
}

//{{.StructName}}Delete 删除
func {{.StructName}}Delete(c *gin.Context) {
	logger.Logger.Debug("In {{.StructName}}Delete")

	err := services.{{.StructName}}DeleteHandle()
	if err != nil {
		logger.Logger.Error(err.Error())
		reply.Response(c, reply.SetCodeMessage(errorUtil.ErrServer.Code, "服务端处理报错"))
		return
	}
	reply.Response(c)
}

//管理员列表
func {{.StructName}}List(c *gin.Context) {
	logger.Logger.Debug("In {{.StructName}}List")

	type ReqParam struct {
		Name     string `form:"name" json:"name" binding:"omitempty"`           //查询名字
		RoleName string `form:"role_name" json:"role_name" binding:"omitempty"` //查询角色
		Page     int    `form:"page" json:"page" binding:"omitempty"`           //当前页
		PerPage  int    `form:"per_page" json:"per_page" binding:"omitempty"`   //每页多少
	}

	var reqParams ReqParam
	err := c.Bind(&reqParams)
	if err != nil {
		logger.Logger.Error(err.Error())
		reply.Response(c, reply.SetCodeMessage(errorUtil.ErrParam.Code, "参数传递不正确"))
		return
	}
	logger.Logger.Debug(fmt.Sprint(reqParams))

	name := strings.TrimSpace(reqParams.Name)
	list, err := services.AdminListHandle(name, reqParams.RoleName, page, perPage)
	if err != nil {
		logger.Logger.Error(err.Error())
		reply.Response(c, reply.SetCodeMessage(errorUtil.ErrServer.Code, "服务端处理报错"))
		return
	}
	reply.Response(c)
}
