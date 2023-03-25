package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var UserScoreLogDao = newUserScoreLogDao()

func newUserScoreLogDao() *userScoreLogDao {
	return &userScoreLogDao{}
}

type userScoreLogDao struct {
}

func (d *userScoreLogDao) Get(db *gorm.DB, id int64) *model.UserScoreLog {
	ret := &model.UserScoreLog{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userScoreLogDao) Take(db *gorm.DB, where ...interface{}) *model.UserScoreLog {
	ret := &model.UserScoreLog{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userScoreLogDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserScoreLog) {
	cnd.Find(db, &list)
	return
}

func (d *userScoreLogDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserScoreLog {
	ret := &model.UserScoreLog{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *userScoreLogDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserScoreLog, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *userScoreLogDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserScoreLog, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserScoreLog{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *userScoreLogDao) Create(db *gorm.DB, t *model.UserScoreLog) (err error) {
	err = db.Create(t).Error
	return
}

func (d *userScoreLogDao) Update(db *gorm.DB, t *model.UserScoreLog) (err error) {
	err = db.Save(t).Error
	return
}

func (d *userScoreLogDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserScoreLog{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *userScoreLogDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserScoreLog{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *userScoreLogDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserScoreLog{}, "id = ?", id)
}
