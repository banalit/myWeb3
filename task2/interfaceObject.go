package task2

import (
	"fmt"
	"math"
)

/*
*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func ShapeTest() {
	var c = Circle{
		Radius: 10,
	}
	fmt.Printf("circle area:%v, perimeter:%v \n", c.Area(), c.Perimeter())

	var r = Rectangle{
		Length: 10,
		Width:  20,
	}
	fmt.Printf("rectangle %v area:%v, perimeter:%v \n", r, r.Area(), r.Perimeter())
}

/**
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
**/

type Person struct {
	Name string
	Age  uint8
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("employid: %s\n", e.EmployeeID)
	fmt.Printf("name:%s\n", e.Name)
	fmt.Printf("age:%d\n", e.Age)
}

func TestEmployee() {
	var employee = Employee{
		Person: Person{
			Name: "joe",
			Age:  10,
		},
		EmployeeID: "11",
	}
	employee.PrintInfo()
}
