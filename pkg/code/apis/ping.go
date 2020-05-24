package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/solate/util/gin/reply"
)

func Ping(c *gin.Context) {
	reply.Response(c, reply.SetData("Pong"))
}
