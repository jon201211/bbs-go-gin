package eventhandler

import (
	"bbs-go/model/constants"
	"bbs-go/service"
	"bbs-go/util/event"
	"reflect"
)

func init() {
	event.RegHandler(reflect.TypeOf(event.TopicDeleteEvent{}), handleTopicDeleteEvent)
}

func handleTopicDeleteEvent(i interface{}) {
	e := i.(event.TopicDeleteEvent)

	// 处理userFeed
	service.UserFeedService.DeleteByDataId(e.TopicId, constants.EntityTopic)

	// 发送消息
	service.MessageService.SendTopicDeleteMsg(e.TopicId, e.DeleteUserId)

	// 操作日志
	service.OperateLogService.AddOperateLog(e.DeleteUserId, constants.OpTypeDelete, constants.EntityTopic,
		e.TopicId, "", nil)
}
