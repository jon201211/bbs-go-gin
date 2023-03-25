package admin

import (
	"bbs-go/controller/base"
	"bbs-go/util/simple/date"
	"strconv"

	"bbs-go/util/simple"

	"bbs-go/model"
	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type TopicNodeController struct {
	base.BaseController
}

func (c *TopicNodeController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.TopicNodeService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *TopicNodeController) AnyList(ctx *gin.Context) {
	list, paging := service.TopicNodeService.FindPageByParams(simple.NewQueryParams(ctx).EqByReq("name").PageByReq().Asc("sort_no").Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}

func (c *TopicNodeController) PostCreate(ctx *gin.Context) {
	t := &model.TopicNode{}
	err := simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t.CreateTime = date.NowTimestamp()
	err = service.TopicNodeService.Create(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *TopicNodeController) PostUpdate(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t := service.TopicNodeService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	err = simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	err = service.TopicNodeService.Update(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *TopicNodeController) GetNodes(ctx *gin.Context) {
	list := service.TopicNodeService.GetNodes()
	c.JsonSuccess(ctx, list)
	return
}
