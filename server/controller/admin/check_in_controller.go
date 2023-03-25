package admin

import (
	"bbs-go/controller/base"
	"bbs-go/model"
	"bbs-go/service"
	"bbs-go/util/simple"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CheckInController struct {
	base.BaseController
}

func (c *CheckInController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.CheckInService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *CheckInController) AnyList(ctx *gin.Context) {
	list, paging := service.CheckInService.FindPageByParams(simple.NewQueryParams(ctx).PageByReq().Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}

func (c *CheckInController) PostCreate(ctx *gin.Context) {
	t := &model.CheckIn{}
	err := simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	err = service.CheckInService.Create(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *CheckInController) PostUpdate(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t := service.CheckInService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	err = simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	err = service.CheckInService.Update(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}
