package admin

import (
	"bbs-go/util/markdown"
	"strconv"

	"bbs-go/controller/base"
	"bbs-go/controller/render"

	"bbs-go/util/simple"

	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	base.BaseController
}

func (c *CommentController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.CommentService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *CommentController) AnyList(ctx *gin.Context) {
	var (
		id         = simple.FormValueInt64Default(ctx, "id", 0)
		userId     = simple.FormValueInt64Default(ctx, "userId", 0)
		entityType = simple.FormValueDefault(ctx, "entityType", "")
		entityId   = simple.FormValueInt64Default(ctx, "entityId", 0)
	)
	params := simple.NewQueryParams(ctx).
		EqByReq("status").
		PageByReq().Desc("id")

	if id > 0 {
		params.Eq("id", id)
	}
	if userId > 0 {
		params.Eq("user_id", userId)
	}
	if simple.IsNotBlank(entityType) && entityId > 0 {
		params.Eq("entity_type", entityType).Eq("entity_id", entityId)
	}

	if id <= 0 && userId <= 0 && (simple.IsBlank(entityType) || entityId <= 0) {
		// c.JsonErrorMsg(ctx, "请输入必要的查询参数。")
		c.JsonSuccess(ctx, nil)
		return
	}

	list, paging := service.CommentService.FindPageByParams(params)

	var results []map[string]interface{}
	for _, comment := range list {
		builder := simple.NewRspBuilderExcludes(comment, "content")

		// 用户
		builder = builder.Put("user", render.BuildUserInfoDefaultIfNull(comment.UserId))

		// 简介
		content := markdown.ToHTML(comment.Content)
		builder.Put("content", content)

		results = append(results, builder.Build())
	}

	c.JsonPageData(ctx, results, paging)
	return
}

func (c *CommentController) PostDeleteBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	if err := service.CommentService.Delete(id); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	} else {
		c.JsonSuccess(ctx, nil)
		return
	}
}
