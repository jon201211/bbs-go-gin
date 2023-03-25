package api

import (
	"bbs-go/model/constants"
	"bbs-go/util/validate"
	"strconv"
	"strings"

	"bbs-go/util/simple"

	"github.com/gin-gonic/gin"

	"bbs-go/cache"
	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/model"
	"bbs-go/service"
)

type UserController struct {
	base.BaseController
}

// 获取当前登录用户
func (c *UserController) GetCurrent(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user != nil {
		c.JsonSuccess(ctx, render.BuildUserProfile(user))
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 用户详情
func (c *UserController) GetBy(ctx *gin.Context) {
	userId := simple.ParamValueInt64Default(ctx, "id", 0)
	user := cache.UserCache.Get(userId)
	if user != nil && user.Status != constants.StatusDeleted {
		c.JsonSuccess(ctx, render.BuildUserProfile(user))
		return
	}
	c.JsonErrorMsg(ctx, "用户不存在")
	return
}

// 修改用户资料
func (c *UserController) PostEditBy(ctx *gin.Context) {
	userId := simple.FormValueInt64Default(ctx, "id", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	if user.Id != userId {
		c.JsonErrorMsg(ctx, "无权限")
		return
	}
	nickname := strings.TrimSpace(simple.FormValue(ctx, "nickname"))
	avatar := strings.TrimSpace(simple.FormValue(ctx, "avatar"))
	homePage := simple.FormValue(ctx, "homePage")
	description := simple.FormValue(ctx, "description")

	if len(nickname) == 0 {
		c.JsonErrorMsg(ctx, "昵称不能为空")
		return
	}
	if len(avatar) == 0 {
		c.JsonErrorMsg(ctx, "头像不能为空")
		return
	}

	if len(homePage) > 0 && validate.IsURL(homePage) != nil {
		c.JsonErrorMsg(ctx, "个人主页地址错误")
		return
	}

	err := service.UserService.Updates(user.Id, map[string]interface{}{
		"nickname":    nickname,
		"avatar":      avatar,
		"home_page":   homePage,
		"description": description,
	})
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 修改头像
func (c *UserController) PostUpdateAvatar(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	avatar := strings.TrimSpace(simple.FormValue(ctx, "avatar"))
	if len(avatar) == 0 {
		c.JsonErrorMsg(ctx, "头像不能为空")
		return
	}
	err := service.UserService.UpdateAvatar(user.Id, avatar)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 设置用户名
func (c *UserController) PostSetUsername(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	username := strings.TrimSpace(simple.FormValue(ctx, "username"))
	err := service.UserService.SetUsername(user.Id, username)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 设置邮箱
func (c *UserController) PostSetEmail(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	email := strings.TrimSpace(simple.FormValue(ctx, "email"))
	err := service.UserService.SetEmail(user.Id, email)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
}

// 设置密码
func (c *UserController) PostSetPassword(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	password := simple.FormValue(ctx, "password")
	rePassword := simple.FormValue(ctx, "rePassword")
	err := service.UserService.SetPassword(user.Id, password, rePassword)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 修改密码
func (c *UserController) PostUpdatePassword(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	var (
		oldPassword = simple.FormValue(ctx, "oldPassword")
		password    = simple.FormValue(ctx, "password")
		rePassword  = simple.FormValue(ctx, "rePassword")
	)
	if err := service.UserService.UpdatePassword(user.Id, oldPassword, password, rePassword); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 设置背景图
func (c *UserController) PostSetBackgroundImage(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	backgroundImage := simple.FormValue(ctx, "backgroundImage")
	if simple.IsBlank(backgroundImage) {
		c.JsonErrorMsg(ctx, "请上传图片")
		return
	}
	if err := service.UserService.UpdateBackgroundImage(user.Id, backgroundImage); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 用户收藏
func (c *UserController) GetFavorites(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	cursor := simple.FormValueInt64Default(ctx, "cursor", 0)

	// 用户必须登录
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}

	// 查询列表
	limit := 20
	var favorites []model.Favorite
	if cursor > 0 {
		favorites = service.FavoriteService.Find(simple.NewSqlCnd().Where("user_id = ? and id < ?",
			user.Id, cursor).Desc("id").Limit(20))
	} else {
		favorites = service.FavoriteService.Find(simple.NewSqlCnd().Where("user_id = ?", user.Id).Desc("id").Limit(limit))
	}

	hasMore := false
	if len(favorites) > 0 {
		cursor = favorites[len(favorites)-1].Id
		hasMore = len(favorites) >= limit
	}

	c.JsonCursorData(ctx, render.BuildFavorites(favorites), strconv.FormatInt(cursor, 10), hasMore)
	return
}

// 获取最近3条未读消息
func (c *UserController) GetMsgrecent(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	var count int64 = 0
	var messages []model.Message
	if user != nil {
		count = service.MessageService.GetUnReadCount(user.Id)
		messages = service.MessageService.Find(simple.NewSqlCnd().Eq("user_id", user.Id).
			Eq("status", constants.MsgStatusUnread).Limit(3).Desc("id"))
	}
	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().Put("count", count).Put("messages", render.BuildMessages(messages)).Build())
	return
}

// 用户消息
func (c *UserController) GetMessages(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	page := simple.FormValueIntDefault(ctx, "page", 1)

	// 用户必须登录
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}

	messages, paging := service.MessageService.FindPageByCnd(simple.NewSqlCnd().
		Eq("user_id", user.Id).
		Page(page, 20).Desc("id"))

	// 全部标记为已读
	service.MessageService.MarkRead(user.Id)

	c.JsonPageData(ctx, render.BuildMessages(messages), paging)
	return
}

// 用户积分记录
func (c *UserController) GetScorelogs(ctx *gin.Context) {
	page := simple.FormValueIntDefault(ctx, "page", 1)
	user := service.UserTokenService.GetCurrent(ctx)
	// 用户必须登录
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}

	logs, paging := service.UserScoreLogService.FindPageByCnd(simple.NewSqlCnd().
		Eq("user_id", user.Id).
		Page(page, 20).Desc("id"))

	c.JsonPageData(ctx, logs, paging)
	return
}

// 积分排行
func (c *UserController) GetScoreRank(ctx *gin.Context) {
	users := cache.UserCache.GetScoreRank()
	var results []*model.UserInfo
	for _, user := range users {
		results = append(results, render.BuildUserInfo(&user))
	}
	c.JsonSuccess(ctx, results)
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
	if days == -1 && !user.HasRole(constants.RoleOwner) {
		c.JsonErrorMsg(ctx, "无永久禁言权限")
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

// PostEmailVerify 请求邮箱验证邮件S
func (c *UserController) PostEmailVerify(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	if err := service.UserService.SendEmailVerifyEmail(user.Id); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// GetEmailVerify 获取邮箱验证码
func (c *UserController) GetEmailVerify(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	token := simple.FormValue(ctx, "token")
	if simple.IsBlank(token) {
		c.JsonErrorMsg(ctx, "非法请求")
		return
	}
	if err := service.UserService.VerifyEmail(user.Id, token); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}
