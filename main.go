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

}
