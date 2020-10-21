package schema

import (
	"../engine"
	"fmt"
	"reflect"
	"strings"
)

// 删表/建表/CRUD
type ORM struct {
	sqlEngine string
	engine    engine.DB
	schema    *Schema
}

func NewORM(sqlEngine string, args string) *ORM {
	return &ORM{
		sqlEngine: sqlEngine,
		engine:    engine.Factory(sqlEngine, args), // 此处可以通过sqlEngine选择不同的数据库引擎，sqlite/mysql
	}
}

// 获取schema，为下一步建表做准备
func (o *ORM) Model(v interface{}) *ORM {
	if o.schema == nil || reflect.TypeOf(v) != reflect.TypeOf(o.schema.Model) {
		o.schema = NewSchema(v)
	}
	return o
}

// 建表
func (o *ORM) CreateTable() error {
	table := o.schema
	var colums []string

	for _, v := range table.fieldMap {
		tmp := fmt.Sprintf("%s %s %s", v.Name, v.Type, v.Tag)
		colums = append(colums, tmp)
	}
	desc := strings.Join(colums, ",")
	tableSql := fmt.Sprintf("Create Table %s (%s)", table.Name, desc)
	o.engine.CreateTable(tableSql)
	fmt.Println(tableSql)
	return nil
}

// 删表
func (o *ORM) DropTable() error {
	sql := fmt.Sprintf("Drop table if exist %s", o.schema.Name)
	fmt.Println(sql)
	return nil
}
