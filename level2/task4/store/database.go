package store

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	GormDB *gorm.DB
}

func GetDB() *Database {
	gormDB := startDB()

	return &Database{GormDB: gormDB}
}

func startDB() *gorm.DB {
	sqlDB, err := sql.Open("mysql", "root:1715rjvxbr7410@tcp(127.0.0.1:6603)/task6")
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return gormDB
}
