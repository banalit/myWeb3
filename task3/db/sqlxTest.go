package task3

import "fmt"

type Employee struct {
	ID         uint64 `gorm:primarykey`
	Name       string
	Department string
	Salary     uint64
}

type Book struct {
	ID     uint64
	Title  string
	Author string
	Price  float64
}

func migrateData() {
	db := getGormDb()
	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&Book{})
	rows, err := db.Model(&Employee{}).Where("1=1 limit 1").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var employee Employee
	for rows.Next() {
		db.ScanRows(rows, &employee)
		return
	}
	employees := []Employee{{Name: "peter", Department: "finace", Salary: 1000},
		{Name: "alice", Department: "finace", Salary: 2000},
		{Name: "apple", Department: "IT", Salary: 3000},
		{Name: "banana", Department: "IT", Salary: 3100},
		{Name: "lucy", Department: "sales", Salary: 4000},
		{Name: "joe", Department: "sales", Salary: 5000}}
	db.Create(employees)
}

func prepareData() {
	migrateData()
	insertBooks()
}
func SqlxTest() {
	prepareData()
	sqlxdb := getSqlxDb()
	var employees []Employee
	sqlxdb.Select(&employees, "select * from employees where department= ?", "IT")
	var emp Employee
	sqlxdb.Get(&emp, "select * from employees where salary = (select max(salary) from employees)")
	fmt.Println("IT employees:", employees)
	fmt.Println("max salary employee:", emp)

	var books []Book
	sqlxdb.Select(&books, "select * from books where price>50")
	fmt.Println("books:", books)

}

func insertBooks() {
	sqlxdb := getSqlxDb()
	var book Book
	err := sqlxdb.Get(&book, "select * from books limit 1")
	if err == nil {
		return
	}
	sqlxdb.MustExec("insert into books ( title, author, price) values(?,?,?)", "happy", "peter", 60)
	sqlxdb.MustExec("insert into books ( title, author, price) values(?,?,?)", "sad", "luke", 12)
	sqlxdb.MustExec("insert into books ( title, author, price) values(?,?,?)", "sweet", "apple", 44.9)
	sqlxdb.MustExec("insert into books ( title, author, price) values(?,?,?)", "bad", "peter", 99.5)

}
