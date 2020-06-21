//%s
func %s(c *gin.Context) {

	reply := new(pro.%s)

	req := new(pro.%s)
	if err := server.BindRequest(c, req); err != nil {
		c.JSON(http.StatusOK, std.BuildAdminResponse(false, std.AdminCodeParamError))
		return
	}
	if err := moduleman.Call(c, moduleman.Request{
		Module: "pro.Pro",
		Func:   "%s",
		Params: req,
	}, reply); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, std.BuildApiResponse(false, err.Code, err.Msg))
		return
	}
	c.JSON(http.StatusOK, std.BuildApiResponse(reply))

}
