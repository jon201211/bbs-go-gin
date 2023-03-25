package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var UserDao = newUserDao()

func newUserDao() *userDao {
	return &userDao{}
}

type userDao struct {
}

func (d *userDao) Get(db *gorm.DB, id int64) *model.User {
	ret := &model.User{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userDao) Take(db *gorm.DB, where ...interface{}) *model.User {
	ret := &model.User{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *userDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.User) {
	cnd.Find(db, &list)
	return
}

func (d *userDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.User {
	ret := &model.User{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *userDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.User, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *userDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.User, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.User{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *userDao) Create(db *gorm.DB, t *model.User) (err error) {
	err = db.Create(t).Error
	return
}

func (d *userDao) Update(db *gorm.DB, t *model.User) (err error) {
	err = db.Save(t).Error
	return
}

func (d *userDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *userDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *userDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.User{}, "id = ?", id)
}

func (d *userDao) GetByEmail(db *gorm.DB, email string) *model.User {
	return d.Take(db, "email = ?", email)
}

func (d *userDao) GetByUsername(db *gorm.DB, username string) *model.User {
	return d.Take(db, "username = ?", username)
}
