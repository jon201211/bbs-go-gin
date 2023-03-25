package api

import (
	"bbs-go/model"
	"bbs-go/model/constants"
	"bbs-go/spam"
	"strconv"

	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"

	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/service"
)

type CommentController struct {
	base.BaseController
}

func (c *CommentController) GetFuck(ctx *gin.Context) {
	go func() {
		users := service.UserService.Find(simple.NewSqlCnd().Eq("forbidden_end_time", -1))
		for _, user := range users {
			// 删除评论
			service.CommentService.ScanByUser(user.Id, func(comments []model.Comment) {
				for _, comment := range comments {
					if comment.Status != constants.StatusDeleted {
						_ = service.CommentService.Delete(comment.Id)
					}
				}
			})
		}
	}()
	c.JsonSuccess(ctx, nil)
	return
}

func (c *CommentController) GetList(ctx *gin.Context) {
	var (
		err        error
		cursor     int64
		entityType string
		entityId   int64
	)
	cursor = simple.FormValueInt64Default(ctx, "cursor", 0)

	if entityType, err = simple.FormValueRequired(ctx, "entityType"); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	if entityId, err = simple.FormValueInt64(ctx, "entityId"); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	comments, cursor, hasMore := service.CommentService.GetComments(entityType, entityId, cursor)
	c.JsonCursorData(ctx, render.BuildComments(comments), strconv.FormatInt(cursor, 10), hasMore)
	return
}

func (c *CommentController) PostCreate(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	form := model.GetCreateCommentForm(ctx)
	if err := spam.CheckComment(user, form); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	comment, err := service.CommentService.Publish(user.Id, form)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	c.JsonSuccess(ctx, render.BuildComment(*comment))
	return
}
