package common

import "bbs-go/util/simple"

var (
	CaptchaError        = simple.NewError(1000, "验证码错误")
	ForbiddenError      = simple.NewError(1001, "已被禁言")
	UserDisabled        = simple.NewError(1002, "账号已禁用")
	InObservationPeriod = simple.NewError(1003, "账号尚在观察期")
	EmailNotVerified    = simple.NewError(1004, "请先验证邮箱")
)
