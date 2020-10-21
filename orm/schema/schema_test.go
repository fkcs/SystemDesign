package schema

import "testing"

func TestORM(t *testing.T) {
	type User struct {
		name string `orm:"PRIMARY KEY"`
		sex  string
		age  int
	}
	session := NewORM("sqlite", "test.db").Model(&User{})
	session.CreateTable()
	session.DropTable()
}
