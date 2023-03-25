package admin

import (
	"bbs-go/controller/base"
	"bbs-go/service"
	"bbs-go/util/simple"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OperateLogController struct {
	base.BaseController
}

func (c *OperateLogController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.OperateLogService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *OperateLogController) AnyList(ctx *gin.Context) {
	list, paging := service.OperateLogService.FindPageByParams(simple.NewQueryParams(ctx).
		EqByReq("user_id").EqByReq("op_type").EqByReq("data_type").EqByReq("data_id").
		PageByReq().Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}
