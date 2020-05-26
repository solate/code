package models

import (
	"github.com/solate/code/pkg/component/orm"
	ormUtil "github.com/solate/util/orm"
)

type {{.StructName}} struct {
	ormUtil.Model
	{{- range $key, $value := .AttributeList}}
	{{$value.Name}} {{$value.Type}} `{{$value.Tag}}` // {{$value.Description}}
    {{- end}}

}

// 数据库表名
func (s *{{.StructName}}) TableName() string {
	return "{{.TableName}}"
}

//新增
//只是新增用这个
func (s *{{.StructName}}) Create() (err error) {
	err = orm.DB.Create(&s).Error
	return
}

//更新一条数据
func (s *{{.StructName}}) Update() error {
	err := orm.DB.Model(&{{.StructName}}{}).Updates(s).Error
	return err
}

func (s *{{.StructName}}) UpdateByConds(id int, conds map[string]interface{}) error {
	err := orm.DB.Model(&{{.StructName}}{}).Where("id = ?", id).Updates(conds).Error
	return err
}

//根据id查找一条数据
func (s *{{.StructName}}) FindById(id int) (err error) {
	err = orm.DB.Where("id = ?", id).First(&s).Error
	return
}

//根据其他条件查询
func (s *{{.StructName}}) Find(conds map[string]interface{}) (err error) {
	err = orm.DB.Where(conds).First(&s).Error
	return
}

//根据名称判断是否唯一
//true 找不到， false 已存在
func (s *{{.StructName}}) RecordNotFound() bool {
	return orm.DB.Where("name = ?", s.Name).First(&s).RecordNotFound()
}

//获得列表,
//conds interface{} where语句参数
//page int  当前页码
//perPage  每页有几个数据,默认15条
func (s *{{.StructName}}) FindList(conds interface{}, page, perPage int) (list []{{.StructName}}, count int, err error) {
	whereDb := orm.DB.Where(conds)
	whereDb.Model(&{{.StructName}}{}).Count(&count) //获得总数
	if perPage > 0 {
		cursor := ormUtil.GetPageStart(page, perPage)
		whereDb = whereDb.Limit(perPage).Offset(cursor)
	}
	err = whereDb.Find(&feeRules).Error
	return
}
