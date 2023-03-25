package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var UserLikeDao = newUserLikeDao()

func newUserLikeDao() *userLikeDao {
	return &userLikeDao{}
}

type userLikeDao struct {
}

func (d *userLikeDao) Get(db *gorm.DB, id int64) *model.UserLike {
	ret := &model.UserLike{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userLikeDao) Take(db *gorm.DB, where ...interface{}) *model.UserLike {
	ret := &model.UserLike{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userLikeDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserLike) {
	cnd.Find(db, &list)
	return
}

func (d *userLikeDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserLike {
	ret := &model.UserLike{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *userLikeDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserLike, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *userLikeDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserLike, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserLike{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *userLikeDao) Create(db *gorm.DB, t *model.UserLike) (err error) {
	err = db.Create(t).Error
	return
}

func (d *userLikeDao) Update(db *gorm.DB, t *model.UserLike) (err error) {
	err = db.Save(t).Error
	return
}

func (d *userLikeDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserLike{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *userLikeDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserLike{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *userLikeDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserLike{}, "id = ?", id)
}
