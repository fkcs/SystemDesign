package engine

import (
	"database/sql"
	"fmt"
	"log"
)

func init() {
	name := "sqlite"
	Register(name, &SQLite{})
}

type SQLite struct {
	db *sql.DB
}

func (s *SQLite) NewDB(arg string) DB {
	db, err := sql.Open("sqlite3", arg)
	if err != nil {
		fmt.Println(err)
	}
	return &SQLite{db: db}
}

func (s *SQLite) CreateTable(sql string) error {
	stat, err := s.db.Prepare(sql)
	if err != nil {
		log.Println(err)
	}
	stat.Exec()
	return nil
}

func (s *SQLite) DropTable(sql string) error {
	return nil
}
