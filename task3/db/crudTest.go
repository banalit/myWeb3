package task3

import (
	"fmt"
	"time"
)

type Student struct {
	ID       uint64 `gorm:"primaeryKey"`
	Name     string
	Age      uint
	Grade    string
	CreateAt time.Time
	UpdateAt time.Time
}

func CrudTest() {
	db := getGormDb()
	student := &Student{
		Name:  "zhangsan",
		Age:   20,
		Grade: "grade 3",
	}
	db.AutoMigrate(student)

	db.Create(student)

	students := []*Student{
		{Name: "peter", Age: 11, Grade: "Grade 2"},
		{Name: "Alice", Age: 20, Grade: "Grade 6"},
		{Name: "Luke", Age: 3, Grade: "Grade 12"},
	}
	db.CreateInBatches(students, 10)

	rows, err := db.Model(&Student{}).Debug().Where("age > 18").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var std Student
	for rows.Next() {
		db.ScanRows(rows, &std)
		fmt.Println("student:", std)
	}

	db.Model(&Student{}).Where("name=?", "zhangsan").Update("grade", "grade 4")
	db.Debug().Where("age < ?", 100).Delete(&Student{})

}
