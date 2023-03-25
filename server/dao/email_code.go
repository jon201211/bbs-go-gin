package dao

import (
	"bbs-go/model"
	"bbs-go/util/simple"

	"gorm.io/gorm"
)

var EmailCodeDao = newEmailCodeDao()

func newEmailCodeDao() *emailCodeDao {
	return &emailCodeDao{}
}

type emailCodeDao struct {
}

func (d *emailCodeDao) Get(db *gorm.DB, id int64) *model.EmailCode {
	ret := &model.EmailCode{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *emailCodeDao) Take(db *gorm.DB, where ...interface{}) *model.EmailCode {
	ret := &model.EmailCode{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *emailCodeDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.EmailCode) {
	cnd.Find(db, &list)
	return
}

func (d *emailCodeDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.EmailCode {
	ret := &model.EmailCode{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *emailCodeDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.EmailCode, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *emailCodeDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.EmailCode, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.EmailCode{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *emailCodeDao) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.EmailCode{})
}

func (d *emailCodeDao) Create(db *gorm.DB, t *model.EmailCode) (err error) {
	err = db.Create(t).Error
	return
}

func (d *emailCodeDao) Update(db *gorm.DB, t *model.EmailCode) (err error) {
	err = db.Save(t).Error
	return
}

func (d *emailCodeDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.EmailCode{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *emailCodeDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.EmailCode{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *emailCodeDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.EmailCode{}, "id = ?", id)
}
