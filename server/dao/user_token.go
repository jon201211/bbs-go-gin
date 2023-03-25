package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var UserTokenDao = newUserTokenDao()

func newUserTokenDao() *userTokenDao {
	return &userTokenDao{}
}

type userTokenDao struct {
}

func (d *userTokenDao) GetByToken(db *gorm.DB, token string) *model.UserToken {
	if len(token) == 0 {
		return nil
	}
	return d.Take(db, "token = ?", token)
}

func (d *userTokenDao) Get(db *gorm.DB, id int64) *model.UserToken {
	ret := &model.UserToken{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userTokenDao) Take(db *gorm.DB, where ...interface{}) *model.UserToken {
	ret := &model.UserToken{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userTokenDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserToken) {
	cnd.Find(db, &list)
	return
}

func (d *userTokenDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserToken {
	ret := &model.UserToken{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *userTokenDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserToken, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *userTokenDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserToken, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserToken{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *userTokenDao) Create(db *gorm.DB, t *model.UserToken) (err error) {
	err = db.Create(t).Error
	return
}

func (d *userTokenDao) Update(db *gorm.DB, t *model.UserToken) (err error) {
	err = db.Save(t).Error
	return
}

func (d *userTokenDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserToken{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *userTokenDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserToken{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *userTokenDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserToken{}, "id = ?", id)
}
