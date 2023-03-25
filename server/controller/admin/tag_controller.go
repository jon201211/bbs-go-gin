package admin

import (
	"bbs-go/model/constants"
	"bbs-go/util/simple/date"
	"strconv"
	"strings"

	"bbs-go/util/simple"

	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/model"
	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	base.BaseController
}

func (c *TagController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.TagService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *TagController) AnyList(ctx *gin.Context) {
	list, paging := service.TagService.FindPageByParams(simple.NewQueryParams(ctx).
		LikeByReq("id").
		LikeByReq("name").
		EqByReq("status").
		PageByReq().Desc("id"))
	c.JsonPageData(ctx, list, paging)
	return
}

func (c *TagController) PostCreate(ctx *gin.Context) {
	t := &model.Tag{}
	err := simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	if len(t.Name) == 0 {
		c.JsonErrorMsg(ctx, "name is required")
		return
	}
	if service.TagService.GetByName(t.Name) != nil {
		c.JsonErrorMsg(ctx, "标签「"+t.Name+"」已存在")
		return
	}

	t.Status = constants.StatusOk
	t.CreateTime = date.NowTimestamp()
	t.UpdateTime = date.NowTimestamp()

	err = service.TagService.Create(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *TagController) PostUpdate(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t := service.TagService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	err = simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	if len(t.Name) == 0 {
		c.JsonErrorMsg(ctx, "name is required")
		return
	}
	if tmp := service.TagService.GetByName(t.Name); tmp != nil && tmp.Id != id {
		c.JsonErrorMsg(ctx, "标签「"+t.Name+"」已存在")
		return
	}

	t.UpdateTime = date.NowTimestamp()
	err = service.TagService.Update(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

// 自动完成
func (c *TagController) GetAutocomplete(ctx *gin.Context) {
	keyword := strings.TrimSpace(ctx.Param("keyword"))
	var tags []model.Tag
	if len(keyword) > 0 {
		tags = service.TagService.Find(simple.NewSqlCnd().Starting("name", keyword).Desc("id"))
	} else {
		tags = service.TagService.Find(simple.NewSqlCnd().Desc("id").Limit(10))
	}
	c.JsonSuccess(ctx, render.BuildTags(tags))
	return
}

// 根据标签编号批量获取
func (c *TagController) GetTags(ctx *gin.Context) {
	tagIds := simple.FormValueInt64Array(ctx, "tagIds")
	var tags *[]model.TagResponse
	if len(tagIds) > 0 {
		tagArr := service.TagService.Find(simple.NewSqlCnd().In("id", tagIds))
		if len(tagArr) > 0 {
			tags = render.BuildTags(tagArr)
		}
	}
	c.JsonSuccess(ctx, tags)
	return
}
