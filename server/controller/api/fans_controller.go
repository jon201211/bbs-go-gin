package api

import (
	"bbs-go/controller/base"
	"bbs-go/controller/render"
	"bbs-go/model"
	"bbs-go/service"
	"strconv"

	"bbs-go/util/simple"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/gin-gonic/gin"
)

type FansController struct {
	base.BaseController
}

func (c *FansController) PostFollow(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}

	otherId := simple.FormValueInt64Default(ctx, "userId", 0)
	if otherId <= 0 {
		c.JsonErrorMsg(ctx, "param: userId required")
		return
	}

	err := service.UserFollowService.Follow(user.Id, otherId)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

func (c *FansController) PostUnfollow(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}

	otherId := simple.FormValueInt64Default(ctx, "userId", 0)
	if otherId <= 0 {
		c.JsonErrorMsg(ctx, "param: userId required")
		return
	}

	err := service.UserFollowService.UnFollow(user.Id, otherId)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

func (c *FansController) GetIsfollowed(ctx *gin.Context) {
	userId := simple.QueryValueInt64Default(ctx, "userId", 0)
	current := service.UserTokenService.GetCurrent(ctx)
	var followed = false
	if current != nil && current.Id != userId {
		followed = service.UserFollowService.IsFollowed(current.Id, userId)
	}
	c.JsonSuccess(ctx, followed)
	return
}

func (c *FansController) GetFans(ctx *gin.Context) {
	userId := simple.FormValueInt64Default(ctx, "userId", 0)
	cursor := simple.FormValueInt64Default(ctx, "cursor", 0)
	userIds, cursor, hasMore := service.UserFollowService.GetFans(userId, cursor, 10)

	current := service.UserTokenService.GetCurrent(ctx)
	var followedSet hashset.Set
	if current != nil {
		followedSet = service.UserFollowService.IsFollowedUsers(current.Id, userIds...)
	}

	var itemList []*model.UserInfo
	for _, id := range userIds {
		item := render.BuildUserInfoDefaultIfNull(id)
		item.Followed = followedSet.Contains(id)
		itemList = append(itemList, item)
	}
	c.JsonCursorData(ctx, itemList, strconv.FormatInt(cursor, 10), hasMore)
	return
}

func (c *FansController) GetFollows(ctx *gin.Context) {
	userId := simple.FormValueInt64Default(ctx, "userId", 0)
	cursor := simple.FormValueInt64Default(ctx, "cursor", 0)
	userIds, cursor, hasMore := service.UserFollowService.GetFollows(userId, cursor, 10)

	current := service.UserTokenService.GetCurrent(ctx)
	var followedSet hashset.Set
	if current != nil {
		if current.Id == userId {
			followedSet = *hashset.New()
			for _, id := range userIds {
				followedSet.Add(id)
			}
		} else {
			followedSet = service.UserFollowService.IsFollowedUsers(current.Id, userIds...)
		}
	}

	var itemList []*model.UserInfo
	for _, id := range userIds {
		item := render.BuildUserInfoDefaultIfNull(id)
		item.Followed = followedSet.Contains(id)
		itemList = append(itemList, item)
	}
	c.JsonCursorData(ctx, itemList, strconv.FormatInt(cursor, 10), hasMore)
	return
}

func (c *FansController) GetRecentFans(ctx *gin.Context) {
	userId := simple.QueryValueInt64Default(ctx, "userId", 0)
	userIds, cursor, hasMore := service.UserFollowService.GetFans(userId, 0, 10)

	current := service.UserTokenService.GetCurrent(ctx)
	var followedSet hashset.Set
	if current != nil {
		followedSet = service.UserFollowService.IsFollowedUsers(current.Id, userIds...)
	}

	var itemList []*model.UserInfo
	for _, id := range userIds {
		item := render.BuildUserInfoDefaultIfNull(id)
		item.Followed = followedSet.Contains(id)
		itemList = append(itemList, item)
	}
	c.JsonCursorData(ctx, itemList, strconv.FormatInt(cursor, 10), hasMore)
	return
}

func (c *FansController) GetRecentFollow(ctx *gin.Context) {
	userId := simple.QueryValueInt64Default(ctx, "userId", 0)
	userIds, cursor, hasMore := service.UserFollowService.GetFollows(userId, 0, 10)

	current := service.UserTokenService.GetCurrent(ctx)
	var followedSet hashset.Set
	if current != nil {
		if current.Id == userId {
			followedSet = *hashset.New()
			for _, id := range userIds {
				followedSet.Add(id)
			}
		} else {
			followedSet = service.UserFollowService.IsFollowedUsers(current.Id, userIds...)
		}
	}

	var itemList []*model.UserInfo
	for _, id := range userIds {
		item := render.BuildUserInfoDefaultIfNull(id)
		item.Followed = followedSet.Contains(id)
		itemList = append(itemList, item)
	}
	c.JsonCursorData(ctx, itemList, strconv.FormatInt(cursor, 10), hasMore)
	return
}
