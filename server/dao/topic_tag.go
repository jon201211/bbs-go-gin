package dao

import (
	"bbs-go/util/simple"
	"bbs-go/util/simple/date"

	"gorm.io/gorm"

	"bbs-go/model"
)

var TopicTagDao = newTopicTagDao()

func newTopicTagDao() *topicTagDao {
	return &topicTagDao{}
}

type topicTagDao struct {
}

func (d *topicTagDao) Get(db *gorm.DB, id int64) *model.TopicTag {
	ret := &model.TopicTag{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *topicTagDao) Take(db *gorm.DB, where ...interface{}) *model.TopicTag {
	ret := &model.TopicTag{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *topicTagDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.TopicTag) {
	cnd.Find(db, &list)
	return
}

func (d *topicTagDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.TopicTag {
	ret := &model.TopicTag{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *topicTagDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.TopicTag, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *topicTagDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.TopicTag, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.TopicTag{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *topicTagDao) Create(db *gorm.DB, t *model.TopicTag) (err error) {
	err = db.Create(t).Error
	return
}

func (d *topicTagDao) Update(db *gorm.DB, t *model.TopicTag) (err error) {
	err = db.Save(t).Error
	return
}

func (d *topicTagDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.TopicTag{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *topicTagDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.TopicTag{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *topicTagDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.TopicTag{}, "id = ?", id)
}

func (d *topicTagDao) AddTopicTags(db *gorm.DB, topicId int64, tagIds []int64) {
	if topicId <= 0 || len(tagIds) == 0 {
		return
	}
	for _, tagId := range tagIds {
		_ = d.Create(db, &model.TopicTag{
			TopicId:    topicId,
			TagId:      tagId,
			CreateTime: date.NowTimestamp(),
		})
	}
}

func (d *topicTagDao) DeleteTopicTags(db *gorm.DB, topicId int64) {
	if topicId <= 0 {
		return
	}
	db.Where("topic_id = ?", topicId).Delete(model.TopicTag{})
}
