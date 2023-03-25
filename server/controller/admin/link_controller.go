package admin

import (
	"bbs-go/controller/base"
	"bbs-go/util/simple/date"
	"bytes"
	"strconv"

	"bbs-go/util/simple"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"

	"bbs-go/model"
	"bbs-go/service"
)

type LinkController struct {
	base.BaseController
}

func (c *LinkController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.LinkService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *LinkController) AnyList(ctx *gin.Context) {
	list, paging := service.LinkService.FindPageByParams(simple.NewQueryParams(ctx).EqByReq("status").LikeByReq("title").LikeByReq("url").PageByReq().Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}

func (c *LinkController) PostCreate(ctx *gin.Context) {
	t := &model.Link{}
	err := simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t.CreateTime = date.NowTimestamp()

	err = service.LinkService.Create(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *LinkController) PostUpdate(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t := service.LinkService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	err = simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	err = service.LinkService.Update(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *LinkController) GetDetect(ctx *gin.Context) {
	url := simple.FormValue(ctx, "url")
	resp, err := resty.New().SetRedirectPolicy(resty.FlexibleRedirectPolicy(3)).R().Get(url)
	if err != nil {
		logrus.Error(err)
		c.JsonSuccess(ctx, nil)
		return
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		logrus.Error(err)
		c.JsonSuccess(ctx, nil)
		return
	}
	title := doc.Find("title").Text()
	description := doc.Find("meta[name=description]").AttrOr("content", "")
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("title", title).Put("description", description).Build())
	return
}
