package dao

import (
	"bbs-go/util/simple"

	"gorm.io/gorm"

	"bbs-go/model"
)

var MessageDao = newMessageDao()

func newMessageDao() *messageDao {
	return &messageDao{}
}

type messageDao struct {
}

func (d *messageDao) Get(db *gorm.DB, id int64) *model.Message {
	ret := &model.Message{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *messageDao) Take(db *gorm.DB, where ...interface{}) *model.Message {
	ret := &model.Message{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *messageDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Message) {
	cnd.Find(db, &list)
	return
}

func (d *messageDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Message {
	ret := &model.Message{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *messageDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Message, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *messageDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Message, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Message{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *messageDao) Create(db *gorm.DB, t *model.Message) (err error) {
	err = db.Create(t).Error
	return
}

func (d *messageDao) Update(db *gorm.DB, t *model.Message) (err error) {
	err = db.Save(t).Error
	return
}

func (d *messageDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Message{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *messageDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Message{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *messageDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Message{}, "id = ?", id)
}
