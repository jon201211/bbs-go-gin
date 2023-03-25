package service

import (
	"bbs-go/dao"
	"bbs-go/model"
	"bbs-go/util/simple"
)

var EmailCodeService = newEmailCodeService()

func newEmailCodeService() *emailCodeService {
	return &emailCodeService{}
}

type emailCodeService struct {
}

func (s *emailCodeService) Get(id int64) *model.EmailCode {
	return dao.EmailCodeDao.Get(simple.DB(), id)
}

func (s *emailCodeService) Take(where ...interface{}) *model.EmailCode {
	return dao.EmailCodeDao.Take(simple.DB(), where...)
}

func (s *emailCodeService) Find(cnd *simple.SqlCnd) []model.EmailCode {
	return dao.EmailCodeDao.Find(simple.DB(), cnd)
}

func (s *emailCodeService) FindOne(cnd *simple.SqlCnd) *model.EmailCode {
	return dao.EmailCodeDao.FindOne(simple.DB(), cnd)
}

func (s *emailCodeService) FindPageByParams(params *simple.QueryParams) (list []model.EmailCode, paging *simple.Paging) {
	return dao.EmailCodeDao.FindPageByParams(simple.DB(), params)
}

func (s *emailCodeService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.EmailCode, paging *simple.Paging) {
	return dao.EmailCodeDao.FindPageByCnd(simple.DB(), cnd)
}

func (s *emailCodeService) Count(cnd *simple.SqlCnd) int64 {
	return dao.EmailCodeDao.Count(simple.DB(), cnd)
}

func (s *emailCodeService) Create(t *model.EmailCode) error {
	return dao.EmailCodeDao.Create(simple.DB(), t)
}

func (s *emailCodeService) Update(t *model.EmailCode) error {
	return dao.EmailCodeDao.Update(simple.DB(), t)
}

func (s *emailCodeService) Updates(id int64, columns map[string]interface{}) error {
	return dao.EmailCodeDao.Updates(simple.DB(), id, columns)
}

func (s *emailCodeService) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.EmailCodeDao.UpdateColumn(simple.DB(), id, name, value)
}

func (s *emailCodeService) Delete(id int64) {
	dao.EmailCodeDao.Delete(simple.DB(), id)
}
