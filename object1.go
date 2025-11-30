//题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64 //计算周长
}
type Rectangle struct {
	Width, Height float64
}

// 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义Circle结构体
type Circle struct {
	Radius float64
}

// 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

//第二题 结构体组合

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person     //嵌入Person结构体
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println("===== 员工信息 =====")
	fmt.Printf("员工ID：%s\n", e.EmployeeID)
	// 直接访问组合的 Person 结构体的字段（字段提升）
	fmt.Printf("姓名：%s\n", e.Name)
	fmt.Printf("年龄：%d\n", e.Age)
}
