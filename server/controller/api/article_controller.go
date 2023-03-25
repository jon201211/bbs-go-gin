package api

import (
	"bbs-go/model/constants"
	"bbs-go/spam"
	"bbs-go/util/urls"
	"math/rand"
	"strconv"

	"bbs-go/util/simple"

	"bbs-go/cache"
	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/model"
	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	base.BaseController
}

// 文章详情
func (c *ArticleController) GetBy(ctx *gin.Context) {
	articleId := simple.ParamValueInt64Default(ctx, "id", 0)
	article := service.ArticleService.Get(articleId)
	if article == nil || article.Status == constants.StatusDeleted {
		c.JsonErrorCode(ctx, 404, "文章不存在")
		return
	}

	user := service.UserTokenService.GetCurrent(ctx)
	if user != nil {
		if article.UserId != user.Id && article.Status == constants.StatusPending {
			c.JsonErrorCode(ctx, 403, "文章审核中")
			return
		}
	} else {
		if article.Status == constants.StatusPending {
			c.JsonErrorCode(ctx, 403, "文章审核中")
			return
		}
	}

	service.ArticleService.IncrViewCount(articleId) // 增加浏览量
	c.JsonSuccess(ctx, render.BuildArticle(article))
	return
}

// PostCreate 发表文章
func (c *ArticleController) PostCreate(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	var (
		tags    = simple.FormValueStringArray(ctx, "tags")
		title   = simple.FormValue(ctx, "title")
		summary = simple.FormValue(ctx, "summary")
		content = simple.FormValue(ctx, "content")
	)
	form := model.CreateArticleForm{
		Title:       title,
		Summary:     summary,
		Content:     content,
		ContentType: constants.ContentTypeMarkdown,
		Tags:        tags,
	}

	if err := spam.CheckArticle(user, form); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	article, err := service.ArticleService.Publish(user.Id, form)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, render.BuildArticle(article))
	return
}

// 编辑时获取详情
func (c *ArticleController) GetEditBy(ctx *gin.Context) {
	articleId := simple.ParamValueInt64Default(ctx, "id", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	article := service.ArticleService.Get(articleId)
	if article == nil || article.Status == constants.StatusDeleted {
		c.JsonErrorMsg(ctx, "话题不存在或已被删除")
		return
	}

	// 非作者、且非管理员
	if article.UserId != user.Id && !user.HasAnyRole(constants.RoleAdmin, constants.RoleOwner) {
		c.JsonErrorMsg(ctx, "无权限")
		return
	}

	tags := service.ArticleService.GetArticleTags(articleId)
	var tagNames []string
	if len(tags) > 0 {
		for _, tag := range tags {
			tagNames = append(tagNames, tag.Name)
		}
	}

	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().
		Put("articleId", article.Id).
		Put("title", article.Title).
		Put("content", article.Content).
		Put("tags", tagNames).
		Build())
}

// 编辑文章
func (c *ArticleController) PostEditBy(ctx *gin.Context) {
	articleId := simple.ParamValueInt64Default(ctx, "id", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	var (
		tags    = simple.FormValueStringArray(ctx, "tags")
		title   = simple.FormValue(ctx, "title")
		content = simple.FormValue(ctx, "content")
	)

	article := service.ArticleService.Get(articleId)
	if article == nil || article.Status == constants.StatusDeleted {
		c.JsonErrorMsg(ctx, "文章不存在")
		return
	}

	// 非作者、且非管理员
	if article.UserId != user.Id && !user.HasAnyRole(constants.RoleAdmin, constants.RoleOwner) {
		c.JsonErrorMsg(ctx, "无权限")
	}

	if err := service.ArticleService.Edit(articleId, tags, title, content); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	// 操作日志
	service.OperateLogService.AddOperateLog(user.Id, constants.OpTypeUpdate, constants.EntityArticle, articleId,
		"", ctx.Request)
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("articleId", article.Id).Build())
	return
}

// 删除文章
func (c *ArticleController) PostDeleteBy(ctx *gin.Context) {
	articleId := simple.ParamValueInt64Default(ctx, "id", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	article := service.ArticleService.Get(articleId)
	if article == nil || article.Status == constants.StatusDeleted {
		c.JsonErrorMsg(ctx, "文章不存在")
		return
	}

	// 非作者、且非管理员
	if article.UserId != user.Id && !user.HasAnyRole(constants.RoleAdmin, constants.RoleOwner) {
		c.JsonErrorMsg(ctx, "无权限")
		return
	}

	if err := service.ArticleService.Delete(articleId); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	// 操作日志
	service.OperateLogService.AddOperateLog(user.Id, constants.OpTypeDelete, constants.EntityArticle, articleId,
		"", ctx.Request)
	c.JsonSuccess(ctx, nil)
	return
}

// 收藏文章
func (c *ArticleController) PostFavoriteBy(ctx *gin.Context) {
	articleId := simple.ParamValueInt64Default(ctx, "id", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	err := service.FavoriteService.AddArticleFavorite(user.Id, articleId)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 文章跳转链接
func (c *ArticleController) GetRedirectBy(ctx *gin.Context) {
	articleId := simple.ParamValueInt64Default(ctx, "id", 0)
	article := service.ArticleService.Get(articleId)
	if article == nil || article.Status != constants.StatusOk {
		c.JsonErrorMsg(ctx, "文章不存在")
		return
	}
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("url", urls.ArticleUrl(articleId)).Build())
}

// 最近文章
func (c *ArticleController) GetRecent(ctx *gin.Context) {
	articles := service.ArticleService.Find(simple.NewSqlCnd().Where("status = ?", constants.StatusOk).Desc("id").Limit(10))
	c.JsonSuccess(ctx, render.BuildSimpleArticles(articles))
	return
}

// 用户文章列表
func (c *ArticleController) GetUserArticles(ctx *gin.Context) {
	userId := simple.QueryValueInt64Default(ctx, "userId", 0)
	if userId <= 0 {
		c.JsonErrorMsg(ctx, "用户不存在")
		return
	}
	cursor := simple.FormValueInt64Default(ctx, "cursor", 0)
	articles, cursor, hasMore := service.ArticleService.GetUserArticles(userId, cursor)
	c.JsonCursorData(ctx, render.BuildSimpleArticles(articles), strconv.FormatInt(cursor, 10), hasMore)
	return
}

// 文章列表
func (c *ArticleController) GetArticles(ctx *gin.Context) {
	cursor := simple.FormValueInt64Default(ctx, "cursor", 0)
	articles, cursor, hasMore := service.ArticleService.GetArticles(cursor)
	c.JsonCursorData(ctx, render.BuildSimpleArticles(articles), strconv.FormatInt(cursor, 10), hasMore)
	return
}

// 标签文章列表
func (c *ArticleController) GetTagArticles(ctx *gin.Context) {
	cursor := simple.FormValueInt64Default(ctx, "cursor", 0)
	tagId := simple.FormValueInt64Default(ctx, "tagId", 0)
	articles, cursor, hasMore := service.ArticleService.GetTagArticles(tagId, cursor)
	c.JsonCursorData(ctx, render.BuildSimpleArticles(articles), strconv.FormatInt(cursor, 10), hasMore)
	return
}

// 用户最新的文章
func (c *ArticleController) GetUserNewestBy(ctx *gin.Context) {
	userId := simple.ParamValueInt64Default(ctx, "userId", 0)
	articles := service.ArticleService.GetUserNewestArticles(userId)
	c.JsonSuccess(ctx, render.BuildSimpleArticles(articles))
	return
}

// 近期文章
func (c *ArticleController) GetNearlyBy(ctx *gin.Context) {
	articleId := simple.ParamValueInt64Default(ctx, "id", 0)
	articles := service.ArticleService.GetNearlyArticles(articleId)
	c.JsonSuccess(ctx, render.BuildSimpleArticles(articles))
	return
}

// 相关文章
func (c *ArticleController) GetRelatedBy(ctx *gin.Context) {
	articleId := simple.ParamValueInt64Default(ctx, "id", 0)
	relatedArticles := service.ArticleService.GetRelatedArticles(articleId)
	c.JsonSuccess(ctx, render.BuildSimpleArticles(relatedArticles))
	return
}

// 推荐
func (c *ArticleController) GetRecommend(ctx *gin.Context) {
	articles := cache.ArticleCache.GetRecommendArticles()
	if len(articles) == 0 {
		c.JsonSuccess(ctx, nil)
		return
	} else {
		dest := make([]model.Article, len(articles))
		perm := rand.Perm(len(articles))
		for i, v := range perm {
			dest[v] = articles[i]
		}
		end := 10
		if end > len(articles) {
			end = len(articles)
		}
		ret := dest[0:end]
		c.JsonSuccess(ctx, render.BuildSimpleArticles(ret))
		return
	}
}

// 最新文章
func (c *ArticleController) GetNewest(ctx *gin.Context) {
	articles := service.ArticleService.Find(simple.NewSqlCnd().Eq("status", constants.StatusOk).Desc("id").Limit(5))
	c.JsonSuccess(ctx, render.BuildSimpleArticles(articles))
	return
}

// 热门文章
func (c *ArticleController) GetHot(ctx *gin.Context) {
	articles := cache.ArticleCache.GetHotArticles()
	c.JsonSuccess(ctx, render.BuildSimpleArticles(articles))
	return
}
