package admin

import (
	"bbs-go/controller/base"
	"bbs-go/model"
	"bbs-go/service"
	"bbs-go/util/simple"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserFollowController struct {
	base.BaseController
}

func (c *UserFollowController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.UserFollowService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *UserFollowController) AnyList(ctx *gin.Context) {
	list, paging := service.UserFollowService.FindPageByParams(simple.NewQueryParams(ctx).PageByReq().Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}

func (c *UserFollowController) PostCreate(ctx *gin.Context) {
	t := &model.UserFollow{}
	err := simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	err = service.UserFollowService.Create(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *UserFollowController) PostUpdate(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t := service.UserFollowService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	err = simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	err = service.UserFollowService.Update(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}
