package engine

import (
	"fmt"
	"reflect"
)

var engines map[string]reflect.Type

func init() {
	engines = make(map[string]reflect.Type)
}

type DB interface {
	NewDB(arg string) DB
	CreateTable(sql string) error
	DropTable(sql string) error
	/*
		Insert(key, value interface{}) error
		Update(key, value interface{}) error
		Delete(key interface{}) error
		Select(key interface{}) error
		Limit() DB
		Where() DB
		OrderBy() DB*/
}

func Register(name string, value interface{}) {
	if _, ok := engines[name]; ok {
		fmt.Println("engines is exist! ", name)
		return
	}
	engines[name] = reflect.TypeOf(value).Elem()
}

func Factory(engine string, arg string) DB {
	if _, ok := engines[engine]; ok {
		c := reflect.New(engines[engine]).Interface().(DB)
		return c.NewDB(arg)
	}
	fmt.Println("invalid engine! ", engine)
	return nil
}
