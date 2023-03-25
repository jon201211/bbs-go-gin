package api

import (
	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"

	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/service"
)

type ProjectController struct {
	base.BaseController
}

func (c *ProjectController) GetBy(ctx *gin.Context) {
	projectId := simple.ParamValueInt64Default(ctx, "id", 0)
	project := service.ProjectService.Get(projectId)
	if project == nil {
		c.JsonErrorMsg(ctx, "项目不存在")
		return
	}
	c.JsonSuccess(ctx, render.BuildProject(project))
	return
}

func (c *ProjectController) GetProjects(ctx *gin.Context) {
	page := simple.FormValueIntDefault(ctx, "page", 1)

	projects, paging := service.ProjectService.FindPageByParams(simple.NewQueryParams(ctx).
		Page(page, 20).Desc("id"))

	c.JsonPageData(ctx, render.BuildSimpleProjects(projects), paging)
	return
}
