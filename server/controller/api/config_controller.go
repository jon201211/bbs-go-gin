package api

import (
	"github.com/gin-gonic/gin"

	"bbs-go/controller/base"
	"bbs-go/service"
)

type ConfigController struct {
	base.BaseController
}

func (c *ConfigController) GetConfigs(ctx *gin.Context) {
	config := service.SysConfigService.GetConfig()
	c.JsonSuccess(ctx, config)
}
