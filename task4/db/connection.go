package task3

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getSqlxSqlliteDb() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", "blogs.db")
	db.SetMaxOpenConns(10)
	return db
}

func getGormSqlliteDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("blogs.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("open sqlite error:", err)
	}
	return db
}
