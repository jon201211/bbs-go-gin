package admin

import (
	"strconv"

	"bbs-go/controller/base"
	"bbs-go/util/simple"

	"bbs-go/model"
	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type TopicTagController struct {
	base.BaseController
}

func (c *TopicTagController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.TopicTagService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *TopicTagController) AnyList(ctx *gin.Context) {
	list, paging := service.TopicTagService.FindPageByParams(simple.NewQueryParams(ctx).PageByReq().Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}

func (c *TopicTagController) PostCreate(ctx *gin.Context) {
	t := &model.TopicTag{}
	simple.ReadForm(ctx, t)

	err := service.TopicTagService.Create(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *TopicTagController) PostUpdate(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t := service.TopicTagService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	simple.ReadForm(ctx, t)

	err = service.TopicTagService.Update(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}
