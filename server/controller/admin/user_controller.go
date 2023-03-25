package admin

import (
	"bbs-go/controller/base"
	"bbs-go/model/constants"
	"strconv"
	"strings"

	"bbs-go/model"

	"bbs-go/util/simple"

	"bbs-go/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.BaseController
}

func (c *UserController) GetSynccount(ctx *gin.Context) {
	go func() {
		service.UserService.SyncUserCount()
	}()
	c.JsonSuccess(ctx, nil)
	return
}

func (c *UserController) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.UserService.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id="+strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, c.buildUserItem(t))
	return
}

func (c *UserController) AnyList(ctx *gin.Context) {
	list, paging := service.UserService.FindPageByParams(simple.NewQueryParams(ctx).EqByReq("id").LikeByReq("nickname").EqByReq("username").PageByReq().Desc("id"))
	var itemList []map[string]interface{}
	for _, user := range list {
		itemList = append(itemList, c.buildUserItem(&user))
	}
	c.JsonPageData(ctx, itemList, paging)
	return
}

func (c *UserController) PostCreate(ctx *gin.Context) {
	username := simple.FormValue(ctx, "username")
	email := simple.FormValue(ctx, "email")
	nickname := simple.FormValue(ctx, "nickname")
	password := simple.FormValue(ctx, "password")

	user, err := service.UserService.SignUp(username, email, nickname, password, password)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, c.buildUserItem(user))
	return
}

func (c *UserController) PostUpdate(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	user := service.UserService.Get(id)
	if user == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	username := simple.FormValue(ctx, "username")
	password := simple.FormValue(ctx, "password")
	nickname := simple.FormValue(ctx, "nickname")
	email := simple.FormValue(ctx, "email")
	roles := simple.FormValueStringArray(ctx, "roles")
	status := simple.FormValueIntDefault(ctx, "status", -1)

	user.Username = simple.SqlNullString(username)
	user.Nickname = nickname
	user.Email = simple.SqlNullString(email)
	user.Roles = strings.Join(roles, ",")
	user.Status = status

	if len(password) > 0 {
		user.Password = simple.EncodePassword(password)
	}

	err = service.UserService.Update(user)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, c.buildUserItem(user))
	return
}

// 禁言
func (c *UserController) PostForbidden(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	if !user.HasAnyRole(constants.RoleOwner, constants.RoleAdmin) {
		c.JsonErrorMsg(ctx, "无权限")
		return
	}
	var (
		userId = simple.FormValueInt64Default(ctx, "userId", 0)
		days   = simple.FormValueIntDefault(ctx, "days", 0)
		reason = simple.FormValue(ctx, "reason")
	)
	if userId < 0 {
		c.JsonErrorMsg(ctx, "请传入：userId")
		return
	}
	if days == 0 {
		service.UserService.RemoveForbidden(user.Id, userId, ctx.Request)
	} else {
		if err := service.UserService.Forbidden(user.Id, userId, days, reason, ctx.Request); err != nil {
			c.JsonErrorMsg(ctx, err.Error())
			return
		}
	}
	c.JsonSuccess(ctx, nil)
	return
}

func (c *UserController) buildUserItem(user *model.User) map[string]interface{} {
	return simple.NewRspBuilder(user).
		Put("roles", user.GetRoles()).
		Put("username", user.Username.String).
		Put("email", user.Email.String).
		Put("score", user.Score).
		Put("forbidden", user.IsForbidden()).
		Build()
}
