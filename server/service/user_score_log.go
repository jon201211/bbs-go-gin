package service

import (
	"bbs-go/util/simple"

	"bbs-go/dao"
	"bbs-go/model"
)

var UserScoreLogService = newUserScoreLogService()

func newUserScoreLogService() *userScoreLogService {
	return &userScoreLogService{}
}

type userScoreLogService struct {
}

func (s *userScoreLogService) Get(id int64) *model.UserScoreLog {
	return dao.UserScoreLogDao.Get(simple.DB(), id)
}

func (s *userScoreLogService) Take(where ...interface{}) *model.UserScoreLog {
	return dao.UserScoreLogDao.Take(simple.DB(), where...)
}

func (s *userScoreLogService) Find(cnd *simple.SqlCnd) []model.UserScoreLog {
	return dao.UserScoreLogDao.Find(simple.DB(), cnd)
}

func (s *userScoreLogService) FindOne(cnd *simple.SqlCnd) *model.UserScoreLog {
	return dao.UserScoreLogDao.FindOne(simple.DB(), cnd)
}

func (s *userScoreLogService) FindPageByParams(params *simple.QueryParams) (list []model.UserScoreLog, paging *simple.Paging) {
	return dao.UserScoreLogDao.FindPageByParams(simple.DB(), params)
}

func (s *userScoreLogService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.UserScoreLog, paging *simple.Paging) {
	return dao.UserScoreLogDao.FindPageByCnd(simple.DB(), cnd)
}

func (s *userScoreLogService) Create(t *model.UserScoreLog) error {
	return dao.UserScoreLogDao.Create(simple.DB(), t)
}

func (s *userScoreLogService) Update(t *model.UserScoreLog) error {
	return dao.UserScoreLogDao.Update(simple.DB(), t)
}

func (s *userScoreLogService) Updates(id int64, columns map[string]interface{}) error {
	return dao.UserScoreLogDao.Updates(simple.DB(), id, columns)
}

func (s *userScoreLogService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.UserScoreLogDao.UpdateColumn(simple.DB(), id, name, value)
}

func (s *userScoreLogService) Delete(id int64) {
	dao.UserScoreLogDao.Delete(simple.DB(), id)
}
