package service

import (
	"bbs-go/model/constants"
	"bbs-go/util/simple"

	"bbs-go/dao"
	"bbs-go/model"
)

var TopicTagService = newTopicTagService()

func newTopicTagService() *topicTagService {
	return &topicTagService{}
}

type topicTagService struct {
}

func (s *topicTagService) Get(id int64) *model.TopicTag {
	return dao.TopicTagDao.Get(simple.DB(), id)
}

func (s *topicTagService) Take(where ...interface{}) *model.TopicTag {
	return dao.TopicTagDao.Take(simple.DB(), where...)
}

func (s *topicTagService) Find(cnd *simple.SqlCnd) []model.TopicTag {
	return dao.TopicTagDao.Find(simple.DB(), cnd)
}

func (s *topicTagService) FindOne(cnd *simple.SqlCnd) *model.TopicTag {
	return dao.TopicTagDao.FindOne(simple.DB(), cnd)
}

func (s *topicTagService) FindPageByParams(params *simple.QueryParams) (list []model.TopicTag, paging *simple.Paging) {
	return dao.TopicTagDao.FindPageByParams(simple.DB(), params)
}

func (s *topicTagService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.TopicTag, paging *simple.Paging) {
	return dao.TopicTagDao.FindPageByCnd(simple.DB(), cnd)
}

func (s *topicTagService) Create(t *model.TopicTag) error {
	return dao.TopicTagDao.Create(simple.DB(), t)
}

func (s *topicTagService) Update(t *model.TopicTag) error {
	return dao.TopicTagDao.Update(simple.DB(), t)
}

func (s *topicTagService) Updates(id int64, columns map[string]interface{}) error {
	return dao.TopicTagDao.Updates(simple.DB(), id, columns)
}

func (s *topicTagService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.TopicTagDao.UpdateColumn(simple.DB(), id, name, value)
}

func (s *topicTagService) DeleteByTopicId(topicId int64) {
	simple.DB().Model(model.TopicTag{}).Where("topic_id = ?", topicId).UpdateColumn("status", constants.StatusDeleted)
}

func (s *topicTagService) UndeleteByTopicId(topicId int64) {
	simple.DB().Model(model.TopicTag{}).Where("topic_id = ?", topicId).UpdateColumn("status", constants.StatusOk)
}
