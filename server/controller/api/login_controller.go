package api

import (
	"bbs-go/controller/base"

	"bbs-go/util/simple"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"

	"bbs-go/service"
	"bbs-go/util/common"
)

type LoginController struct {
	base.BaseController
}

// 注册
func (c *LoginController) PostSignup(ctx *gin.Context) {
	var (
		captchaId   = ctx.PostForm("captchaId")
		captchaCode = ctx.PostForm("captchaCode")
		email       = ctx.PostForm("email")
		username    = ctx.PostForm("username")
		password    = ctx.PostForm("password")
		rePassword  = ctx.PostForm("rePassword")
		nickname    = ctx.PostForm("nickname")
		ref         = simple.FormValue(ctx, "ref")
	)
	loginMethod := service.SysConfigService.GetLoginMethod()
	if !loginMethod.Password {
		c.JsonErrorMsg(ctx, "账号密码登录/注册已禁用")
		return
	}
	if !captcha.VerifyString(captchaId, captchaCode) {
		c.JsonErrorMsg(ctx, common.CaptchaError.Error())
		return
	}
	user, err := service.UserService.SignUp(username, email, nickname, password, rePassword)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.BuildLoginSuccess(ctx, user, ref)
	return
}

// 用户名密码登录
func (c *LoginController) PostSignin(ctx *gin.Context) {
	var (
		captchaId   = ctx.PostForm("captchaId")
		captchaCode = ctx.PostForm("captchaCode")
		username    = ctx.PostForm("username")
		password    = ctx.PostForm("password")
		ref         = simple.FormValue(ctx, "ref")
	)
	// loginMethod := service.SysConfigService.GetLoginMethod()
	// if !loginMethod.Password {
	// 	c.JsonErrorMsg(ctx, "账号密码登录/注册已禁用")
	//  return
	// }
	if !captcha.VerifyString(captchaId, captchaCode) {
		c.JsonErrorMsg(ctx, common.CaptchaError.Error())
		return
	}
	user, err := service.UserService.SignIn(username, password)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.BuildLoginSuccess(ctx, user, ref)
	return
}

// 退出登录
func (c *LoginController) GetSignout(ctx *gin.Context) {
	err := service.UserTokenService.Signout(ctx)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}
