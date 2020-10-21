package schema

import (
	"fmt"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

// 将类型转为字符串
func DataType(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.String:
		return "text"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.Array, reflect.Slice:
		return "blob"
	}
	return fmt.Sprintf("invalid type %s,%s", value.Type().Name(), value.Kind())
}

// 定义每个字段，在表中是一列
type Field struct {
	Name string // 字段名
	Type string // 类型，可以通过dataType转化得到
	Tag  string // 约束条件
}

// 表信息
type Schema struct {
	Model    interface{}       // 表示要映射的对象
	Name     string            // 表名
	fieldMap map[string]*Field // 字段
}

// 解析结构体生成Schema
func NewSchema(value interface{}) *Schema {
	module := reflect.Indirect(reflect.ValueOf(value)).Type() //reflect.TypeOf(value)
	schema := &Schema{
		Model:    value,
		Name:     module.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < module.NumField(); i++ {
		v := module.Field(i)
		field := &Field{
			Name: v.Name,
			Type: DataType(reflect.Indirect(reflect.New(v.Type))),
		}
		if k, ok := v.Tag.Lookup("orm"); ok {
			field.Tag = k
		}
		schema.fieldMap[v.Name] = field
	}
	return schema
}

// 通过字段获取每一列具体信息
func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name]
}
