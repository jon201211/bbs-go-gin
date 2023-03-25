package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var ThirdAccountDao = newThirdAccountDao()

func newThirdAccountDao() *thirdAccountDao {
	return &thirdAccountDao{}
}

type thirdAccountDao struct {
}

func (d *thirdAccountDao) Get(db *gorm.DB, id int64) *model.ThirdAccount {
	ret := &model.ThirdAccount{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *thirdAccountDao) Take(db *gorm.DB, where ...interface{}) *model.ThirdAccount {
	ret := &model.ThirdAccount{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *thirdAccountDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.ThirdAccount) {
	cnd.Find(db, &list)
	return
}

func (d *thirdAccountDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.ThirdAccount {
	ret := &model.ThirdAccount{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *thirdAccountDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.ThirdAccount, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *thirdAccountDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.ThirdAccount, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.ThirdAccount{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *thirdAccountDao) Create(db *gorm.DB, t *model.ThirdAccount) (err error) {
	err = db.Create(t).Error
	return
}

func (d *thirdAccountDao) Update(db *gorm.DB, t *model.ThirdAccount) (err error) {
	err = db.Save(t).Error
	return
}

func (d *thirdAccountDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.ThirdAccount{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *thirdAccountDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.ThirdAccount{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *thirdAccountDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.ThirdAccount{}, "id = ?", id)
}
