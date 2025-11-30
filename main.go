package main

import "fmt"

func main() {
	// 示例：创建一个整数并通过指针传递给 addTen 函数
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

}
