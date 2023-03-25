package simple

import "html/template"

var daoTmpl = template.Must(template.New("dao").Parse(`package dao

import (
	"{{.PkgName}}/model"
	"bbs-go/util/simple"
	"gorm.io/gorm"
)

var {{.Name}}Dao = new{{.Name}}Dao()

func new{{.Name}}Dao() *{{.CamelName}}Dao {
	return &{{.CamelName}}Dao{}
}

type {{.CamelName}}Dao struct {
}

func (d *{{.CamelName}}Dao) Get(db *gorm.DB, id int64) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (d *{{.CamelName}}Dao) Take(db *gorm.DB, where ...interface{}) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (d *{{.CamelName}}Dao) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.{{.Name}}) {
	cnd.Find(db, &list)
	return
}

func (d *{{.CamelName}}Dao) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (d *{{.CamelName}}Dao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.{{.Name}}, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (d *{{.CamelName}}Dao) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.{{.Name}}, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.{{.Name}}{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (d *{{.CamelName}}Dao) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.{{.Name}}{})
}

func (d *{{.CamelName}}Dao) Create(db *gorm.DB, t *model.{{.Name}}) (err error) {
	err = db.Create(t).Error
	return
}

func (d *{{.CamelName}}Dao) Update(db *gorm.DB, t *model.{{.Name}}) (err error) {
	err = db.Save(t).Error
	return
}

func (d *{{.CamelName}}Dao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (d *{{.CamelName}}Dao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (d *{{.CamelName}}Dao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.{{.Name}}{}, "id = ?", id)
}

`))

var serviceTmpl = template.Must(template.New("service").Parse(`package service

import (
	"{{.PkgName}}/model"
	"{{.PkgName}}/dao"
	"bbs-go/util/simple"
)

var {{.Name}}Service = new{{.Name}}Service()

func new{{.Name}}Service() *{{.CamelName}}Service {
	return &{{.CamelName}}Service {}
}

type {{.CamelName}}Service struct {
}

func (s *{{.CamelName}}Service) Get(id int64) *model.{{.Name}} {
	return dao.{{.Name}}Dao.Get(simple.DB(), id)
}

func (s *{{.CamelName}}Service) Take(where ...interface{}) *model.{{.Name}} {
	return dao.{{.Name}}Dao.Take(simple.DB(), where...)
}

func (s *{{.CamelName}}Service) Find(cnd *simple.SqlCnd) []model.{{.Name}} {
	return dao.{{.Name}}Dao.Find(simple.DB(), cnd)
}

func (s *{{.CamelName}}Service) FindOne(cnd *simple.SqlCnd) *model.{{.Name}} {
	return dao.{{.Name}}Dao.FindOne(simple.DB(), cnd)
}

func (s *{{.CamelName}}Service) FindPageByParams(params *simple.QueryParams) (list []model.{{.Name}}, paging *simple.Paging) {
	return dao.{{.Name}}Dao.FindPageByParams(simple.DB(), params)
}

func (s *{{.CamelName}}Service) FindPageByCnd(cnd *simple.SqlCnd) (list []model.{{.Name}}, paging *simple.Paging) {
	return dao.{{.Name}}Dao.FindPageByCnd(simple.DB(), cnd)
}

func (s *{{.CamelName}}Service) Count(cnd *simple.SqlCnd) int64 {
	return dao.{{.Name}}Dao.Count(simple.DB(), cnd)
}

func (s *{{.CamelName}}Service) Create(t *model.{{.Name}}) error {
	return dao.{{.Name}}Dao.Create(simple.DB(), t)
}

func (s *{{.CamelName}}Service) Update(t *model.{{.Name}}) error {
	return dao.{{.Name}}Dao.Update(simple.DB(), t)
}

func (s *{{.CamelName}}Service) Updates(id int64, columns map[string]interface{}) error {
	return dao.{{.Name}}Dao.Updates(simple.DB(), id, columns)
}

func (s *{{.CamelName}}Service) UpdateColumn(id int64, name string, value interface{}) error {
	return dao.{{.Name}}Dao.UpdateColumn(simple.DB(), id, name, value)
}

func (s *{{.CamelName}}Service) Delete(id int64) {
	dao.{{.Name}}Dao.Delete(simple.DB(), id)
}

`))

var controllerTmpl = template.Must(template.New("controller").Parse(`package admin

import (
	"{{.PkgName}}/model"
	"{{.PkgName}}/service"
	"bbs-go/util/simple"
	"bbs-go/controller/base"
	"github.com/gin-gonic/gin"
	"strconv"
)

type {{.Name}}Controller struct {
	base.BaseController
}

func (c *{{.Name}}Controller) GetBy(ctx *gin.Context) {
	id := simple.ParamValueInt64Default(ctx, "id", 0)
	t := service.{{.Name}}Service.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "Not found, id=" + strconv.FormatInt(id, 10))
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

func (c *{{.Name}}Controller) AnyList(ctx *gin.Context) {
	list, paging := service.{{.Name}}Service.FindPageByParams(simple.NewQueryParams(ctx).PageByReq().Desc("id"))
	c.JsonPageData(ctx, list,  paging)
	return
}

func (c *{{.Name}}Controller) PostCreate(ctx *gin.Context) {
	t := &model.{{.Name}}{}
	err := simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	err = service.{{.Name}}Service.Create(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
}

func (c *{{.Name}}Controller) PostUpdate(ctx *gin.Context) {
	id, err := simple.FormValueInt64(ctx, "id")
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	t := service.{{.Name}}Service.Get(id)
	if t == nil {
		c.JsonErrorMsg(ctx, "entity not found")
		return
	}

	err = simple.ReadForm(ctx, t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}

	err = service.{{.Name}}Service.Update(t)
	if err != nil {
		c.JsonErrorMsg(ctx, err.Error())
		return
	}
	c.JsonSuccess(ctx, t)
	return
}

`))

var viewIndexTmpl = template.Must(template.New("index.vue").Parse(`
<template>
    <section class="page-container">
        <!--工具条-->
        <el-col :span="24" class="toolbar">
            <el-form :inline="true" :model="filters">
                <el-form-item>
                    <el-input v-model="filters.name" placeholder="名称"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" v-on:click="list">查询</el-button>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="handleAdd">新增</el-button>
                </el-form-item>
            </el-form>
        </el-col>

        <!--列表-->
        <el-table :data="results" highlight-current-row border v-loading="listLoading"
                  style="width: 100%;" @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column prop="id" label="编号"></el-table-column>
            {{range .Fields}}
			<el-table-column prop="{{.CamelName}}" label="{{.CamelName}}"></el-table-column>
			{{end}}
            <el-table-column label="操作" width="150">
                <template slot-scope="scope">
                    <el-button size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                </template>
            </el-table-column>
        </el-table>

        <!--工具条-->
        <el-col :span="24" class="toolbar">
            <el-pagination layout="total, sizes, prev, pager, next, jumper" :page-sizes="[20, 50, 100, 300]"
                           @current-change="handlePageChange"
                           @size-change="handleLimitChange"
                           :current-page="page.page"
                           :page-size="page.limit"
                           :total="page.total"
                           style="float:right;">
            </el-pagination>
        </el-col>

        <!--新增界面-->
        <el-dialog title="新增" :visible.sync="addFormVisible" :close-on-click-modal="false">
            <el-form :model="addForm" label-width="80px" ref="addForm">
                {{range .Fields}}
				<el-form-item label="{{.CamelName}}">
					<el-input v-model="addForm.{{.CamelName}}"></el-input>
				</el-form-item>
                {{end}}
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click.native="addFormVisible = false">取消</el-button>
                <el-button type="primary" @click.native="addSubmit" :loading="addLoading">提交</el-button>
            </div>
        </el-dialog>

        <!--编辑界面-->
        <el-dialog title="编辑" :visible.sync="editFormVisible" :close-on-click-modal="false">
            <el-form :model="editForm" label-width="80px" ref="editForm">
                <el-input v-model="editForm.id" type="hidden"></el-input>
                {{range .Fields}}
				<el-form-item label="{{.CamelName}}">
					<el-input v-model="editForm.{{.CamelName}}"></el-input>
				</el-form-item>
                {{end}}
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click.native="editFormVisible = false">取消</el-button>
                <el-button type="primary" @click.native="editSubmit" :loading="editLoading">提交</el-button>
            </div>
        </el-dialog>
    </section>
</template>

<script>
  export default {
    data() {
      return {
        results: [],
        listLoading: false,
        page: {},
        filters: {},
        selectedRows: [],

        addForm: {},
        addFormVisible: false,
        addLoading: false,

        editForm: {},
        editFormVisible: false,
        editLoading: false,
      }
    },
    mounted() {
      this.list();
    },
    methods: {
		async list() {
		  const params = Object.assign(this.filters, {
			page: this.page.page,
			limit: this.page.limit,
		  })
		  try {
			const data = await this.$axios.post('/api/admin/{{.KebabName}}/list', params)
			this.results = data.results
			this.page = data.page
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e || e.message })
		  } finally {
			this.listLoading = false
		  }
		},
		async handlePageChange(val) {
		  this.page.page = val
		  await this.list()
		},
		async handleLimitChange(val) {
		  this.page.limit = val
		  await this.list()
		},
		handleAdd() {
		  this.addForm = {
			name: '',
			description: '',
		  }
		  this.addFormVisible = true
		},
		async addSubmit() {
		  try {
			await this.$axios.post('/api/admin/{{.KebabName}}/create', this.addForm)
			this.$message({ message: '提交成功', type: 'success' })
			this.addFormVisible = false
			await this.list()
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e || e.message })
		  }
		},
		async handleEdit(index, row) {
		  try {
			const data = await this.$axios.get('/api/admin/{{.KebabName}}/' + row.id)
			this.editForm = Object.assign({}, data)
			this.editFormVisible = true
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e || e.message })
		  }
		},
		async editSubmit() {
		  try {
			await this.$axios.post('/api/admin/{{.KebabName}}/update', this.editForm)
			await this.list()
			this.editFormVisible = false
		  } catch (e) {
			this.$notify.error({ title: '错误', message: e || e.message })
		  }
		},
	
		handleSelectionChange(val) {
		  this.selectedRows = val
		},
    }
  }
</script>

<style lang="scss" scoped>

</style>

`))
