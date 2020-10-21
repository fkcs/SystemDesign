package engine

import (
	"database/sql"
	"fmt"
	"log"
)

func init() {
	name := "mysql"
	Register(name, &Mysql{})
}

type Mysql struct {
	db *sql.DB
}

func (s *Mysql) NewDB(arg string) DB {
	db, err := sql.Open("mysql", arg)
	if err != nil {
		fmt.Println(err)
	}
	return &Mysql{db: db}
}

func (s *Mysql) CreateTable(sql string) error {
	stat, err := s.db.Prepare(sql)
	if err != nil {
		log.Println(err)
	}
	stat.Exec()
	return nil
}

func (s *Mysql) DropTable(sql string) error {
	return nil
}
