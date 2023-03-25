package api

import (
	"bbs-go/model/constants"
	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"

	"bbs-go/cache"
	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/service"
)

type TagController struct {
	base.BaseController
}

// 标签详情
func (c *TagController) GetBy(ctx *gin.Context) {
	tagId := simple.ParamValueInt64Default(ctx, "id", 0)
	tag := cache.TagCache.Get(tagId)
	if tag == nil {
		c.JsonErrorMsg(ctx, "标签不存在")
		return
	}
	c.JsonSuccess(ctx, render.BuildTag(tag))
	return
}

// 标签列表
func (c *TagController) GetTags(ctx *gin.Context) {
	page := simple.FormValueIntDefault(ctx, "page", 1)
	tags, paging := service.TagService.FindPageByCnd(simple.NewSqlCnd().
		Eq("status", constants.StatusOk).
		Page(page, 200).Desc("id"))

	c.JsonPageData(ctx, render.BuildTags(tags), paging)
	return
}

// 标签自动完成
func (c *TagController) PostAutocomplete(ctx *gin.Context) {
	input := simple.FormValue(ctx, "input")
	tags := service.TagService.Autocomplete(input)
	c.JsonSuccess(ctx, tags)
	return
}
