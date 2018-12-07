package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func genSchema() {
	db.AutoMigrate(&Todo{})
}

func dropTables() {
	db.Exec("DROP TABLE todos;")
}

func migrateDB() {
	dropTables()
	genSchema()
	seedTodos()
}

func seedTodos() {
	db.Create(&Todo{Name: "run 30 min", Due: time.Now()})
	db.Create(&Todo{Name: "buy coffee", Due: time.Now().AddDate(0, 0, 1)})
	db.Create(&Todo{Name: "write a blog post", Due: time.Now().AddDate(0, 1, 0)})
	db.Create(&Todo{Name: "find a new job", Due: time.Now().AddDate(1, 0, 0)})
}

func setupDB() (*gorm.DB, error) {
	DBMS := "mysql"

	CONNECT := "root@/go_todo_app?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)
	return db, err
}
