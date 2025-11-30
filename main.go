package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 创建一个整数并通过指针传递给 addTen 函数
	//测试指针部分
	fmt.Println("Pointer Function Test:")
	val := 7
	fmt.Println("Before:", val)
	addTen(&val)
	fmt.Println("After:", val)
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("修改前的切片：%v\n", nums) // 输出：修改前的切片：[1 2 3 4 5]

	// 2. 调用doubleEach函数，传入切片的地址（&nums获取切片指针）
	MultiplyByTwo(&nums)

	// 3. 输出修改后的原始切片
	fmt.Printf("修改后的切片：%v\n", nums) // 输出：修改后的切片：[2 4 6 8 10]
	fmt.Println("---------")
	//测试协程
	fmt.Println("Goroutine Function Test:")
	limit := 10
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printOddNumbers(limit)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		printEvenNumbers(limit)
	}()
	wg.Wait()
	fmt.Println("所有协程执行完毕") // 创建调度器实例
	scheduler := &Scheduler{}

	// 添加示例任务1：模拟耗时500ms的任务
	scheduler.AddTask(Task{
		ID: "task-1",
		Func: func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("任务1执行完成")
		},
	})

	// 添加示例任务2：模拟耗时300ms的任务
	scheduler.AddTask(Task{
		ID: "task-2",
		Func: func() {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("任务2执行完成")
		},
	})

	// 添加示例任务3：模拟耗时800ms的任务
	scheduler.AddTask(Task{
		ID: "task-3",
		Func: func() {
			time.Sleep(800 * time.Millisecond)
			fmt.Println("任务3执行完成")
		},
	})

	// 执行所有任务并获取结果
	results := scheduler.Run()

	// 输出任务执行统计
	fmt.Println("\n===== 任务执行时间统计 =====")
	for _, res := range results {
		fmt.Printf("任务[%s] 执行耗时: %v\n", res.TaskID, res.Elapsed)
	}
	// 输出示例结束
	fmt.Println("===== 任务调度示例结束 =====")
	//测试接口部分
	fmt.Println("\nInterface Function Test:")
	// 创建 Rectangle 实例（长5，宽3）
	rect := Rectangle{Width: 5, Height: 3}
	// 创建 Circle 实例（半径2）
	circle := Circle{Radius: 2}

	// 调用长方形的方法并打印结果
	fmt.Println("===== 长方形 =====")
	fmt.Printf("长：%.2f, 宽：%.2f\n", rect.Width, rect.Height)
	fmt.Printf("面积：%.2f\n", rect.Area())      // 输出：15.00
	fmt.Printf("周长：%.2f\n", rect.Perimeter()) // 输出：16.00

	// 调用圆形的方法并打印结果
	fmt.Println("\n===== 圆形 =====")
	fmt.Printf("半径：%.2f\n", circle.Radius)
	fmt.Printf("面积：%.2f\n", circle.Area())      // 输出：12.57（π取3.14159...）
	fmt.Printf("周长：%.2f\n", circle.Perimeter()) // 输出：12.57

	//测试结构体组合部分
	emp1 := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "EMP001",
	}

	// 方式2：简化初始化（匿名字段可直接嵌套赋值）
	emp2 := Employee{
		Person: Person{
			Name: "李四",
			Age:  28,
		},
		EmployeeID: "EMP002",
	}

	// 方式3：先创建实例，再逐个赋值
	var emp3 Employee
	emp3.Name = "王五" // 直接赋值 Person 的 Name 字段（字段提升）
	emp3.Age = 35    // 直接赋值 Person 的 Age 字段
	emp3.EmployeeID = "EMP003"

	// 5. 调用 PrintInfo() 方法输出员工信息
	emp1.PrintInfo()
	fmt.Println() // 空行分隔
	emp2.PrintInfo()
	fmt.Println()
	emp3.PrintInfo()
}
