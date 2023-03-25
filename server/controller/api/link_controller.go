package api

import (
	"bbs-go/controller/base"
	"bbs-go/model/constants"
	"bbs-go/util/simple"

	"bbs-go/model"
	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type LinkController struct {
	base.BaseController
}

func (c *LinkController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	link := service.LinkService.Get(id)
	if link == nil || link.Status == constants.StatusDeleted {
		c.JsonErrorMsg(ctx, "数据不存在")
		return
	}
	c.JsonSuccess(ctx, c.buildLink(*link))
	return
}

// 列表
func (c *LinkController) GetLinks(ctx *gin.Context) {
	page := simple.FormValueIntDefault(ctx, "page", 1)

	links, paging := service.LinkService.FindPageByCnd(simple.NewSqlCnd().
		Eq("status", constants.StatusOk).Page(page, 20).Asc("id"))

	var itemList []map[string]interface{}
	for _, v := range links {
		itemList = append(itemList, c.buildLink(v))
	}
	c.JsonPageData(ctx, itemList, paging)
	return
}

// 前10个链接
func (c *LinkController) GetToplinks(ctx *gin.Context) {
	links := service.LinkService.Find(simple.NewSqlCnd().
		Eq("status", constants.StatusOk).Limit(10).Asc("id"))

	var itemList []map[string]interface{}
	for _, v := range links {
		itemList = append(itemList, c.buildLink(v))
	}
	c.JsonSuccess(ctx, itemList)
	return
}

func (c *LinkController) buildLink(link model.Link) map[string]interface{} {
	return map[string]interface{}{
		"linkId":     link.Id,
		"url":        link.Url,
		"title":      link.Title,
		"summary":    link.Summary,
		"logo":       link.Logo,
		"createTime": link.CreateTime,
	}
}
