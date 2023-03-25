package admin

import (
	"strconv"

	"bbs-go/util/simple"

	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type TopicController struct {
	base.BaseController
}

func (c *TopicController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.TopicService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *TopicController) AnyList(ctx *gin.Context) {
	list, paging := service.TopicService.FindPageByParams(simple.NewQueryParams(ctx).
		EqByReq("id").EqByReq("user_id").EqByReq("status").EqByReq("recommend").LikeByReq("title").PageByReq().Desc("id"))

	var results []map[string]interface{}
	for _, topic := range list {
		item := render.BuildSimpleTopic(&topic)
		builder := simple.NewRspBuilder(item)
		builder.Put("status", topic.Status)
		results = append(results, builder.Build())
	}

	c.JsonPageData(ctx, results, paging)
	return
}

// 推荐
func (c *TopicController) PostRecommend(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	err = service.TopicService.SetRecommend(id, true)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 取消推荐
func (c *TopicController) DeleteRecommend(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	err = service.TopicService.SetRecommend(id, false)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

func (c *TopicController) PostDelete(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	err = service.TopicService.Delete(id, user.Id, ctx.Request)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

func (c *TopicController) PostUndelete(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	err = service.TopicService.Undelete(id)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}
