package engine

import (
	"database/sql"
	"fmt"
	"log"
)

type Mysql struct {
	db *sql.DB
}

func NewMysql(arg string) *Mysql {
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
