package api

import (
	"bbs-go/cache"
	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/service"
	"bbs-go/util/simple"
	"bbs-go/util/simple/date"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckinController struct {
	base.BaseController
}

// PostCheckin 签到
func (c *CheckinController) PostCheckin(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	err := service.CheckInService.CheckIn(user.Id)
	if err == nil {
		c.JsonSuccess(ctx, nil)
		return
	} else {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
}

// GetCheckin 获取签到信息
func (c *CheckinController) GetCheckin(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonSuccess(ctx, nil)
		return
	}
	checkIn := service.CheckInService.GetByUserId(user.Id)
	if checkIn != nil {
		today := date.GetDay(time.Now())
		c.JsonSuccess(ctx, simple.NewRspBuilder(checkIn).
			Put("checkIn", checkIn.LatestDayName == today). // 今日是否已签到
			Build())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// GetRank 获取当天签到排行榜（最早签到的排在最前面）
func (c *CheckinController) GetRank(ctx *gin.Context) {
	list := cache.UserCache.GetCheckInRank()
	var itemList []map[string]interface{}
	for _, checkIn := range list {
		itemList = append(itemList, simple.NewRspBuilder(checkIn).
			Put("user", render.BuildUserInfoDefaultIfNull(checkIn.UserId)).
			Build())
	}
	c.JsonSuccess(ctx, itemList)
	return
}
