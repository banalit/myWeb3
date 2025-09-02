package main

import (
	"fmt"

	task3 "github.com/luke/web3Learn/task3/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MYSQL_IP = "47.106.246.115"
const MYSQL_USER = "metanode"
const MYSQL_PWD = "metanodeLuke2025"
const MYSQL_DB = "metanode"

func Connect() *gorm.DB {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:13306)/%s?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PWD, MYSQL_IP, MYSQL_DB)
	db, err := gorm.Open(mysql.Open(connectStr))
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := Connect()
	// task3.CrudTest(db)
	task3.TransactionTest(db)

}
