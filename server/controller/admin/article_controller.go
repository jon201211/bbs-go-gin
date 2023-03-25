package admin

import (
	"bbs-go/model/constants"
	"bbs-go/util/simple/date"
	"bbs-go/util/sitemap"
	"strconv"

	"bbs-go/model"

	"bbs-go/util/simple"

	"bbs-go/cache"
	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/service"
	"bbs-go/util/common"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	base.BaseController
}

func (c *ArticleController) GetSitemap(ctx *gin.Context) {
	go func() {
		sitemap.Generate()
	}()
	c.JsonSuccess(ctx, nil)
}

///{id}
func (c *ArticleController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.ArticleService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
}

func (c *ArticleController) AnyList(ctx *gin.Context) {
	var (
		id     = simple.FormValueInt64Default(ctx, "id", 0)
		userId = simple.FormValueInt64Default(ctx, "userId", 0)
	)
	params := simple.NewQueryParams(ctx)
	if id > 0 {
		params.Eq("id", id)
	}
	if userId > 0 {
		params.Eq("user_id", userId)
	}
	params.EqByReq("status").EqByReq("title").PageByReq().Desc("id")

	if id <= 0 && userId <= 0 {
		c.JsonErrorMsg(ctx, "请指定查询的【文章编号】或【作者编号】")
		return
	}
	list, paging := service.ArticleService.FindPageByParams(params)
	results := c.buildArticles(list)
	c.JsonPageData(ctx, results, paging)
}

// GetRecent 展示最近一页数据
func (c *ArticleController) GetRecent(ctx *gin.Context) {
	params := simple.NewQueryParams(ctx).EqByReq("id").EqByReq("user_id").EqByReq("status").Desc("id").Limit(20)
	list := service.ArticleService.Find(&params.SqlCnd)
	results := c.buildArticles(list)
	c.JsonSuccess(ctx, results)
}

// 构建文章列表返回数据
func (c *ArticleController) buildArticles(articles []model.Article) []map[string]interface{} {
	var results []map[string]interface{}
	for _, article := range articles {
		builder := simple.NewRspBuilderExcludes(article, "content")

		// 用户
		builder = builder.Put("user", render.BuildUserInfoDefaultIfNull(article.UserId))

		// 简介
		builder.Put("summary", common.GetSummary(article.ContentType, article.Content))

		// 标签
		tagIds := cache.ArticleTagCache.Get(article.Id)
		tags := cache.TagCache.GetList(tagIds)
		builder.Put("tags", render.BuildTags(tags))

		results = append(results, builder.Build())
	}
	return results
}

func (c *ArticleController) PostUpdate(ctx *gin.Context) {
	id := simple.FormValueInt64Default(ctx, "id", 0)
	if id <= 0 {
		c.JsonErrorMsg(ctx, "id is required")
		return
	}
	t := service.ArticleService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	if err := simple.ReadForm(ctx, t); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	// 数据校验
	if len(t.Title) == 0 {
		c.JsonErrorMsg(ctx, "标题不能为空")
		return
	}
	if len(t.Content) == 0 {
		c.JsonErrorMsg(ctx, "内容不能为空")
		return
	}
	if len(t.ContentType) == 0 {
		c.JsonErrorMsg(ctx, "请选择内容格式")
		return
	}

	t.UpdateTime = date.NowTimestamp()
	err := service.ArticleService.Update(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	c.JsonSuccess(ctx, t)
}

func (c *ArticleController) GetTags(ctx *gin.Context) {
	articleId := simple.FormValueInt64Default(ctx, "articleId", 0)
	tags := service.ArticleService.GetArticleTags(articleId)
	c.JsonSuccess(ctx, render.BuildTags(tags))
	return
}

func (c *ArticleController) PutTags(ctx *gin.Context) {
	var (
		articleId = simple.FormValueInt64Default(ctx, "articleId", 0)
		tags      = simple.FormValueStringArray(ctx, "tags")
	)
	service.ArticleService.PutTags(articleId, tags)
	c.JsonSuccess(ctx, render.BuildTags(service.ArticleService.GetArticleTags(articleId)))
	return
}

func (c *ArticleController) PostDelete(ctx *gin.Context) {
	id := simple.FormValueInt64Default(ctx, "id", 0)
	if id <= 0 {
		c.JsonErrorMsg(ctx, "id is required")
		return
	}
	err := service.ArticleService.Delete(id)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

func (c *ArticleController) PostPending(ctx *gin.Context) {
	id := simple.FormValueInt64Default(ctx, "id", 0)
	if id <= 0 {
		c.JsonErrorMsg(ctx, "id is required")
		return
	}
	err := service.ArticleService.UpdateColumn(id, "status", constants.StatusOk)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
	}
	c.JsonSuccess(ctx, nil)
	return
}
