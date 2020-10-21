package engine

import "fmt"

type DB interface {
	CreateTable(sql string) error
	DropTable(sql string) error
	/*Insert(key, value interface{}) error
	Update(key, value interface{}) error
	Delete(key interface{}) error
	Select(key interface{}) error
	Limit() DB
	Where() DB
	OrderBy() DB*/
}

func Factory(engine string, arg string) DB {
	switch engine {
	case "sqlite":
		return NewSQLite(arg)
	case "mysql":
		return NewMysql(arg)
	default:
		fmt.Println("invalid engine: ", engine)
	}
	return nil
}
