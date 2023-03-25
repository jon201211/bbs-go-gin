package admin

import (
	"os"
	"runtime"

	"bbs-go/controller/base"
	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"
)

type CommonController struct {
	base.BaseController
}

func (c *CommonController) GetSysteminfo(ctx *gin.Context) {
	hostname, _ := os.Hostname()
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().
		Put("os", runtime.GOOS).
		Put("arch", runtime.GOARCH).
		Put("numCpu", runtime.NumCPU()).
		Put("goversion", runtime.Version()).
		Put("hostname", hostname).
		Build())
}
