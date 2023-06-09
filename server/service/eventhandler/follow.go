package eventhandler

import (
	"bbs-go/model"
	"bbs-go/model/constants"
	"bbs-go/service"
	"bbs-go/util/event"
	"reflect"
)

func init() {
	event.RegHandler(reflect.TypeOf(event.FollowEvent{}), handleFollowEvent)
}

func handleFollowEvent(i interface{}) {
	e := i.(event.FollowEvent)

	// 将该用户下的帖子添加到信息流
	service.TopicService.ScanByUser(e.OtherId, func(topics []model.Topic) {
		for _, topic := range topics {
			if topic.Status != constants.StatusOk {
				continue
			}
			_ = service.UserFeedService.Create(&model.UserFeed{
				UserId:     e.UserId,
				DataType:   constants.EntityTopic,
				DataId:     topic.Id,
				AuthorId:   topic.UserId,
				CreateTime: topic.CreateTime,
			})
		}
	})
}
