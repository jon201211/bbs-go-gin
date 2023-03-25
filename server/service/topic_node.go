package service

import (
	"bbs-go/model/constants"
	"bbs-go/util/simple"

	"bbs-go/dao"
	"bbs-go/model"
)

var TopicNodeService = newTopicNodeService()

func newTopicNodeService() *topicNodeService {
	return &topicNodeService{}
}

type topicNodeService struct {
}

func (s *topicNodeService) Get(id int64) *model.TopicNode {
	return dao.TopicNodeDao.Get(simple.DB(), id)
}

func (s *topicNodeService) Take(where ...interface{}) *model.TopicNode {
	return dao.TopicNodeDao.Take(simple.DB(), where...)
}

func (s *topicNodeService) Find(cnd *simple.SqlCnd) []model.TopicNode {
	return dao.TopicNodeDao.Find(simple.DB(), cnd)
}

func (s *topicNodeService) FindOne(cnd *simple.SqlCnd) *model.TopicNode {
	return dao.TopicNodeDao.FindOne(simple.DB(), cnd)
}

func (s *topicNodeService) FindPageByParams(params *simple.QueryParams) (list []model.TopicNode, paging *simple.Paging) {
	return dao.TopicNodeDao.FindPageByParams(simple.DB(), params)
}

func (s *topicNodeService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.TopicNode, paging *simple.Paging) {
	return dao.TopicNodeDao.FindPageByCnd(simple.DB(), cnd)
}

func (s *topicNodeService) Create(t *model.TopicNode) error {
	return dao.TopicNodeDao.Create(simple.DB(), t)
}

func (s *topicNodeService) Update(t *model.TopicNode) error {
	return dao.TopicNodeDao.Update(simple.DB(), t)
}

func (s *topicNodeService) Updates(id int64, columns map[string]interface{}) error {
	return dao.TopicNodeDao.Updates(simple.DB(), id, columns)
}

func (s *topicNodeService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.TopicNodeDao.UpdateColumn(simple.DB(), id, name, value)
}

func (s *topicNodeService) Delete(id int64) {
	dao.TopicNodeDao.Delete(simple.DB(), id)
}

func (s *topicNodeService) GetNodes() []model.TopicNode {
	return dao.TopicNodeDao.Find(simple.DB(), simple.NewSqlCnd().Eq("status", constants.StatusOk).Asc("sort_no").Desc("id"))
}
