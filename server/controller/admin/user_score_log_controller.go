package admin

import (
	"strconv"

	"bbs-go/util/simple"

	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type UserScoreLogController struct {
	base.BaseController
}

func (c *UserScoreLogController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.UserScoreLogService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
}

func (c *UserScoreLogController) AnyList(ctx *gin.Context) {
	list, paging := service.UserScoreLogService.FindPageByParams(simple.NewQueryParams(ctx).
		EqByReq("user_id").EqByReq("source_type").EqByReq("source_id").EqByReq("type").PageByReq().Desc("id"))

	var results []map[string]interface{}
	for _, userScoreLog := range list {
		user := render.BuildUserInfoDefaultIfNull(userScoreLog.UserId)
		item := simple.NewRspBuilder(userScoreLog).Put("user", user).Build()
		results = append(results, item)
	}

	c.JsonPageData(ctx, results, paging)
}
