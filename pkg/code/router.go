package code

import (
	"github.com/gin-gonic/gin"
	"github.com/solate/code/pkg/code/apis"
)

// 路由表
const (
	prefix = ""
	ping   = "/ping"
)

func Router(r *gin.Engine) {
	v1 := r.Group(prefix)
	{
		v1.GET(ping, apis.Ping)
	}
}
