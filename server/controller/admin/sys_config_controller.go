package admin

import (
	"strconv"

	"bbs-go/controller/base"
	"bbs-go/util/simple"

	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type SysConfigController struct {
	base.BaseController
}

func (c *SysConfigController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.SysConfigService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *SysConfigController) AnyList(ctx *gin.Context) {
	list, paging := service.SysConfigService.FindPageByParams(simple.NewQueryParams(ctx).PageByReq().Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}

func (c *SysConfigController) GetAll(ctx *gin.Context) {
	config := service.SysConfigService.GetConfig()
	c.JsonSuccess(ctx, config)
	return
}

func (c *SysConfigController) PostSave(ctx *gin.Context) {
	config := simple.FormValue(ctx, "config")
	if err := service.SysConfigService.SetAll(config); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}
