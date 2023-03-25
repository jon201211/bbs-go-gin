package dao

import (
	"bbs-go/model"
	"bbs-go/util/simple"

	"gorm.io/gorm"
)

var UserFollowDao = newUserFollowDao()

func newUserFollowDao() *userFollowDao {
	return &userFollowDao{}
}

type userFollowDao struct {
}

func (d *userFollowDao) Get(db *gorm.DB, id int64) *model.UserFollow {
	ret := &model.UserFollow{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userFollowDao) Take(db *gorm.DB, where ...interface{}) *model.UserFollow {
	ret := &model.UserFollow{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userFollowDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserFollow) {
	cnd.Find(db, &list)
	return
}

func (d *userFollowDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserFollow {
	ret := &model.UserFollow{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *userFollowDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserFollow, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *userFollowDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserFollow, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserFollow{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *userFollowDao) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.UserFollow{})
}

func (d *userFollowDao) Create(db *gorm.DB, t *model.UserFollow) (err error) {
	err = db.Create(t).Error
	return
}

func (d *userFollowDao) Update(db *gorm.DB, t *model.UserFollow) (err error) {
	err = db.Save(t).Error
	return
}

func (d *userFollowDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserFollow{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *userFollowDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserFollow{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *userFollowDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserFollow{}, "id = ?", id)
}
