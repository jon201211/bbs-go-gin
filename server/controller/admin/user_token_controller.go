package admin

import (
	"strconv"

	"bbs-go/controller/base"
	"bbs-go/util/simple"

	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type UserTokenController struct {
	base.BaseController
}

func (c *UserTokenController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.UserTokenService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *UserTokenController) AnyList(ctx *gin.Context) {
	list, paging := service.UserTokenService.FindPageByParams(simple.NewQueryParams(ctx).PageByReq().Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}
