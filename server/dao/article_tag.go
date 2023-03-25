package dao

import (
	"bbs-go/util/simple"
	"bbs-go/util/simple/date"

	"gorm.io/gorm"

	"bbs-go/model"
)

var ArticleTagDao = newArticleTagDao()

func newArticleTagDao() *articleTagDao {
	return &articleTagDao{}
}

type articleTagDao struct {
}

func (d *articleTagDao) Get(db *gorm.DB, id int64) *model.ArticleTag {
	ret := &model.ArticleTag{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *articleTagDao) Take(db *gorm.DB, where ...interface{}) *model.ArticleTag {
	ret := &model.ArticleTag{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *articleTagDao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.ArticleTag) {
	cnd.Find(db, &list)
	return
}

func (d *articleTagDao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.ArticleTag {
	ret := &model.ArticleTag{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *articleTagDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.ArticleTag, paging *simple.Paging) {
	return d.FindPageByCnd(db, &params.SqlCnd)
}

func (d *articleTagDao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.ArticleTag, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.ArticleTag{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *articleTagDao) Create(db *gorm.DB, t *model.ArticleTag) (err error) {
	err = db.Create(t).Error
	return
}

func (d *articleTagDao) Update(db *gorm.DB, t *model.ArticleTag) (err error) {
	err = db.Save(t).Error
	return
}

func (d *articleTagDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.ArticleTag{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *articleTagDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.ArticleTag{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *articleTagDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.ArticleTag{}, "id = ?", id)
}

func (d *articleTagDao) AddArticleTags(db *gorm.DB, articleId int64, tagIds []int64) {
	if articleId <= 0 || len(tagIds) == 0 {
		return
	}

	for _, tagId := range tagIds {
		_ = d.Create(db, &model.ArticleTag{
			ArticleId:  articleId,
			TagId:      tagId,
			CreateTime: date.NowTimestamp(),
		})
	}
}

func (d *articleTagDao) DeleteArticleTags(db *gorm.DB, articleId int64) {
	if articleId <= 0 {
		return
	}
	db.Where("article_id = ?", articleId).Delete(model.ArticleTag{})
}

func (d *articleTagDao) DeleteArticleTag(db *gorm.DB, articleId, tagId int64) {
	if articleId <= 0 {
		return
	}
	db.Where("article_id = ? and tag_id = ?", articleId, tagId).Delete(model.ArticleTag{})
}

func (d *articleTagDao) FindByArticleId(db *gorm.DB, articleId int64) []model.ArticleTag {
	return d.Find(db, simple.NewSqlCnd().Where("article_id = ?", articleId))
}
