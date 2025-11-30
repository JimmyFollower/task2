package main

import (
	"fmt"
	"sync"
	"time"
)

// 打印奇函数
func printOddNumbers(limit int) {
	for i := 1; i <= limit; i += 2 {
		fmt.Printf("%d ", i)
	}
}

// 打印偶函数
func printEvenNumbers(limit int) {
	for i := 2; i <= limit; i += 2 {
		fmt.Printf("%d ", i)
	}
}

// 设计任务
type Task struct {
	ID   string
	Func func()
}

type TaskResult struct {
	TaskID  string        // 任务ID
	Elapsed time.Duration // 执行耗时
}
type Scheduler struct {
	tasks   []Task       //待执行任务列表
	results []TaskResult //任务执行结果
	mu      sync.Mutex   //互斥锁
	wg      sync.WaitGroup
}

// AddTask 向调度器添加任务
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// Run 执行所有任务并统计时间
func (s *Scheduler) Run() []TaskResult {
	// 遍历任务，启动协程执行
	for _, task := range s.tasks {
		s.wg.Add(1)
		// 启动协程执行单个任务
		go func(t Task) {
			defer s.wg.Done() // 任务完成后减少计数

			// 记录任务开始时间
			start := time.Now()
			// 执行任务
			t.Func()
			// 计算执行耗时
			elapsed := time.Since(start)

			// 并发安全地写入结果
			s.mu.Lock()
			s.results = append(s.results, TaskResult{
				TaskID:  t.ID,
				Elapsed: elapsed,
			})
			s.mu.Unlock()
		}(task) // 传入当前任务（避免循环变量捕获问题）
	}

	// 等待所有任务执行完成
	s.wg.Wait()
	return s.results
}
