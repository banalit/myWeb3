package task3

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const MYSQL_IP = ""
const MYSQL_USER = ""
const MYSQL_PWD = ""
const MYSQL_DB = ""

func getSqlxDb() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:13306)/%s?parseTime=true&loc=Local", MYSQL_USER, MYSQL_PWD, MYSQL_IP, MYSQL_DB)
	db := sqlx.MustConnect("mysql", dsn)
	db.SetMaxOpenConns(10)
	return db
}

func getGormDb() *gorm.DB {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:13306)/%s?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PWD, MYSQL_IP, MYSQL_DB)
	db, err := gorm.Open(mysql.Open(connectStr))
	if err != nil {
		panic(err)
	}
	return db
}

func getSqlxSqlliteDb() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", "metanode.db")
	db.SetMaxOpenConns(10)
	return db
}

func getGormSqlliteDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("metanode.db"), &gorm.Config{})
	if err != nil {
		log.Println("open sqlite error:", err)
	}
	return db
}
