package api

import (
	"bbs-go/controller/base"
	"bbs-go/model/constants"
	"bbs-go/util/collect"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"

	"bbs-go/service"
)

type SpiderController struct {
	base.BaseController
}

// 微信采集发布接口
func (c *SpiderController) PostWxPublish(ctx *gin.Context) {
	err := c.checkToken(ctx)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	article := &collect.WxArticle{}
	err = simple.ReadForm(ctx, article)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t, err := collect.NewWxbotApi().Publish(article)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("id", t.Id).Build())
	return
}

// 采集发布
func (c *SpiderController) PostArticlePublish(ctx *gin.Context) {
	err := c.checkToken(ctx)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	article := &collect.Article{}
	err = simple.ReadForm(ctx, article)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	articleId, err := collect.NewSpiderApi().Publish(article)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("id", articleId).Build())
	return
}

func (c *SpiderController) PostCommentPublish(ctx *gin.Context) {
	err := c.checkToken(ctx)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	comment := &collect.Comment{}
	err = simple.ReadForm(ctx, comment)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	commentId, err := collect.NewSpiderApi().PublishComment(comment)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("id", commentId).Build())
	return
}

func (c *SpiderController) PostProjectPublish(ctx *gin.Context) {
	err := c.checkToken(ctx)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	var (
		userIdStr   = simple.FormValue(ctx, "userId")
		userId, _   = strconv.ParseInt(userIdStr, 10, 64)
		name        = simple.FormValue(ctx, "name")
		title       = simple.FormValue(ctx, "title")
		logo        = simple.FormValue(ctx, "logo")
		url         = simple.FormValue(ctx, "url")
		docUrl      = simple.FormValue(ctx, "docUrl")
		downloadUrl = simple.FormValue(ctx, "downloadUrl")
		content     = simple.FormValue(ctx, "content")
		contentType = simple.FormValue(ctx, "contentType")
	)

	if len(name) == 0 || len(title) == 0 || len(content) == 0 {
		c.JsonErrorMsg(ctx, "数据不完善...")
		return
	}

	temp := service.ProjectService.FindOne(simple.NewSqlCnd().Eq("name", name))
	if temp != nil {
		c.JsonErrorMsg(ctx, "项目已经存在："+name)
		return
	}

	if len(contentType) == 0 {
		contentType = constants.ContentTypeHtml
	}

	p, err := service.ProjectService.Publish(userId, name, title, logo, url, docUrl, downloadUrl,
		contentType, content)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("id", p.Id).Build())
	return
}

func (c *SpiderController) checkToken(ctx *gin.Context) error {
	token := simple.FormValue(ctx, "token")
	data, err := ioutil.ReadFile("/data/publish_token")
	if err != nil {
		return err
	}
	token2 := strings.TrimSpace(string(data))
	if token != token2 {
		return errors.New("token invalidate")
	}
	return nil
}
