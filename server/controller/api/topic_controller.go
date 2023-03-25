package api

import (
	"bbs-go/model/constants"
	"bbs-go/spam"
	"math/rand"
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

type TopicController struct {
	base.BaseController
}

// 节点
func (c *TopicController) GetNodes(ctx *gin.Context) {
	nodes := service.TopicNodeService.GetNodes()
	c.JsonSuccess(ctx, render.BuildNodes(nodes))
	return
}

// 节点信息
func (c *TopicController) GetNode(ctx *gin.Context) {
	nodeId := simple.QueryValueInt64Default(ctx, "nodeId", 0)
	node := service.TopicNodeService.Get(nodeId)
	c.JsonSuccess(ctx, render.BuildNode(node))
	return
}

// 发表帖子
func (c *TopicController) PostCreate(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	form := model.GetCreateTopicForm(ctx)

	if err := spam.CheckTopic(user, form); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	topic, err := service.TopicService.Publish(user.Id, form)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, render.BuildSimpleTopic(topic))
	return
}

// 编辑时获取详情
func (c *TopicController) GetEditBy(ctx *gin.Context) {
	topicId := simple.ParamValueInt64Default(ctx, "id", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	topic := service.TopicService.Get(topicId)
	if topic == nil || topic.Status != constants.StatusOk {
		c.JsonErrorMsg(ctx, "话题不存在或已被删除")
		return
	}
	if topic.Type != constants.TopicTypeTopic {
		c.JsonErrorMsg(ctx, "当前类型帖子不支持修改")
		return
	}

	// 非作者、且非管理员
	if topic.UserId != user.Id && !user.HasAnyRole(constants.RoleAdmin, constants.RoleOwner) {
		c.JsonErrorMsg(ctx, "无权限")
		return
	}

	tags := service.TopicService.GetTopicTags(topicId)
	var tagNames []string
	if len(tags) > 0 {
		for _, tag := range tags {
			tagNames = append(tagNames, tag.Name)
		}
	}

	c.JsonSuccess(ctx, simple.NewEmptyRspBuilder().
		Put("topicId", topic.Id).
		Put("nodeId", topic.NodeId).
		Put("title", topic.Title).
		Put("content", topic.Content).
		Put("tags", tagNames).
		Build())
}

// 编辑帖子
func (c *TopicController) PostEditBy(ctx *gin.Context) {
	topicId := simple.ParamValueInt64Default(ctx, "id", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	topic := service.TopicService.Get(topicId)
	if topic == nil || topic.Status != constants.StatusOk {
		c.JsonErrorMsg(ctx, "话题不存在或已被删除")
		return
	}

	// 非作者、且非管理员
	if topic.UserId != user.Id && !user.HasAnyRole(constants.RoleAdmin, constants.RoleOwner) {
		c.JsonErrorMsg(ctx, "无权限")
		return
	}

	nodeId := simple.FormValueInt64Default(ctx, "nodeId", 0)
	title := strings.TrimSpace(simple.FormValue(ctx, "title"))
	content := strings.TrimSpace(simple.FormValue(ctx, "content"))
	tags := simple.FormValueStringArray(ctx, "tags")

	err := service.TopicService.Edit(topicId, nodeId, tags, title, content)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	// 操作日志
	service.OperateLogService.AddOperateLog(user.Id, constants.OpTypeUpdate, constants.EntityTopic, topicId,
		"", ctx.Request)
	c.JsonSuccess(ctx, render.BuildSimpleTopic(topic))
	return
}

// 删除帖子
func (c *TopicController) PostDeleteBy(ctx *gin.Context) {
	topicId := simple.ParamValueInt64Default(ctx, "id", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	if err := service.UserService.CheckPostStatus(user); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	topic := service.TopicService.Get(topicId)
	if topic == nil || topic.Status != constants.StatusOk {
		c.JsonSuccess(ctx, nil)
		return
	}

	// 非作者、且非管理员
	if topic.UserId != user.Id && !user.HasAnyRole(constants.RoleAdmin, constants.RoleOwner) {
		c.JsonErrorMsg(ctx, "无权限")
		return
	}

	if err := service.TopicService.Delete(topicId, user.Id, ctx.Request); err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// PostRecommendBy 设为推荐
func (c *TopicController) PostRecommendBy(ctx *gin.Context) {
	topicId := simple.ParamValueInt64Default(ctx, "id", 0)
	recommend, err := simple.FormValueBool(ctx, "recommend")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	if !user.HasAnyRole(constants.RoleOwner, constants.RoleAdmin) {
		c.JsonErrorMsg(ctx, "无权限")
		return
	}

	err = service.TopicService.SetRecommend(topicId, recommend)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 帖子详情
func (c *TopicController) GetBy(ctx *gin.Context) {
	topicId := simple.ParamValueInt64Default(ctx, "id", 0)
	topic := service.TopicService.Get(topicId)
	if topic == nil || topic.Status != constants.StatusOk {
		c.JsonErrorMsg(ctx, "主题不存在")
		return
	}
	service.TopicService.IncrViewCount(topicId) // 增加浏览量
	c.JsonSuccess(ctx, render.BuildTopic(topic))
	return
}

// 点赞
func (c *TopicController) PostLikeBy(ctx *gin.Context) {
	topicId := simple.ParamValueInt64Default(ctx, "id", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	err := service.UserLikeService.TopicLike(user.Id, topicId)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 点赞用户
func (c *TopicController) GetRecentlikesBy(ctx *gin.Context) {
	topicId := simple.ParamValueInt64Default(ctx, "id", 0)
	likes := service.UserLikeService.Recent(constants.EntityTopic, topicId, 5)
	var users []model.UserInfo
	for _, like := range likes {
		userInfo := render.BuildUserInfoDefaultIfNull(like.UserId)
		if userInfo != nil {
			users = append(users, *userInfo)
		}
	}
	c.JsonSuccess(ctx, users)
	return
}

// 最新帖子
func (c *TopicController) GetRecent(ctx *gin.Context) {
	user := service.UserTokenService.GetCurrent(ctx)
	topics := service.TopicService.Find(simple.NewSqlCnd().Where("status = ?", constants.StatusOk).Desc("id").Limit(10))
	c.JsonSuccess(ctx, render.BuildSimpleTopics(topics, user))
	return
}

// 用户帖子列表
func (c *TopicController) GetUserTopics(ctx *gin.Context) {
	userId, err := simple.QueryValueInt64(ctx, "userId")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	cursor := simple.FormValueInt64Default(ctx, "cursor", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	topics, cursor, hasMore := service.TopicService.GetUserTopics(userId, cursor)
	c.JsonCursorData(ctx, render.BuildSimpleTopics(topics, user), strconv.FormatInt(cursor, 10), hasMore)
	return
}

// 帖子列表
func (c *TopicController) GetTopics(ctx *gin.Context) {
	var (
		cursor       = simple.FormValueInt64Default(ctx, "cursor", 0)
		nodeId       = simple.FormValueInt64Default(ctx, "nodeId", 0)
		recommend, _ = simple.FormValueBool(ctx, "recommend")
		user         = service.UserTokenService.GetCurrent(ctx)
	)
	topics, cursor, hasMore := service.TopicService.GetTopics(nodeId, cursor, recommend)
	c.JsonCursorData(ctx, render.BuildSimpleTopics(topics, user), strconv.FormatInt(cursor, 10), hasMore)
	return
}

// 标签帖子列表
func (c *TopicController) GetTagTopics(ctx *gin.Context) {
	var (
		cursor     = simple.FormValueInt64Default(ctx, "cursor", 0)
		tagId, err = simple.FormValueInt64(ctx, "tagId")
		user       = service.UserTokenService.GetCurrent(ctx)
	)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	topics, cursor, hasMore := service.TopicService.GetTagTopics(tagId, cursor)
	c.JsonCursorData(ctx, render.BuildSimpleTopics(topics, user), strconv.FormatInt(cursor, 10), hasMore)
	return
}

// 收藏
func (c *TopicController) GetFavoriteBy(ctx *gin.Context) {
	topicId := simple.ParamValueInt64Default(ctx, "id", 0)
	user := service.UserTokenService.GetCurrent(ctx)
	if user == nil {
		c.JsonErrorMsg(ctx, simple.ErrorNotLogin.Error())
		return
	}
	err := service.FavoriteService.AddTopicFavorite(user.Id, topicId)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, nil)
	return
}

// 推荐话题列表（目前逻辑为取最近50条数据随机展示）
func (c *TopicController) GetRecommend(ctx *gin.Context) {
	topics := cache.TopicCache.GetRecommendTopics()
	if len(topics) == 0 {
		c.JsonSuccess(ctx, nil)
		return
	} else {
		dest := make([]model.Topic, len(topics))
		perm := rand.Perm(len(topics))
		for i, v := range perm {
			dest[v] = topics[i]
		}
		end := 10
		if end > len(topics) {
			end = len(topics)
		}
		ret := dest[0:end]
		c.JsonSuccess(ctx, render.BuildSimpleTopics(ret, nil))
		return
	}
}

// 最新话题
func (c *TopicController) GetNewest(ctx *gin.Context) {
	topics := service.TopicService.Find(simple.NewSqlCnd().Eq("status", constants.StatusOk).Desc("id").Limit(6))
	c.JsonSuccess(ctx, render.BuildSimpleTopics(topics, nil))
	return
}
