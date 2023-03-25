package dao

import (
	"bbs-go/model"
	"bbs-go/util/simple"

	"gorm.io/gorm"
)

var UserFeedDao = newUserFeedDao()

func newUserFeedDao() *userFeedDao {
	return &userFeedDao{}
}

type userFeedDao struct {
}

func (d *userFeedDao) Get(db *gorm.DB, id int64) *model.UserFeed {
	ret := &model.UserFeed{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userFeedDao) Take(db *gorm.DB, where ...interface{}) *model.UserFeed {
	ret := &model.UserFeed{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userFeedDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserFeed) {
	cnd.Find(db, &list)
	return
}

func (d *userFeedDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserFeed {
	ret := &model.UserFeed{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *userFeedDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserFeed, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *userFeedDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserFeed, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserFeed{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *userFeedDao) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.UserFeed{})
}

func (d *userFeedDao) Create(db *gorm.DB, t *model.UserFeed) (err error) {
	err = db.Create(t).Error
	return
}

func (d *userFeedDao) Update(db *gorm.DB, t *model.UserFeed) (err error) {
	err = db.Save(t).Error
	return
}

func (d *userFeedDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserFeed{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *userFeedDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserFeed{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *userFeedDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserFeed{}, "id = ?", id)
}
