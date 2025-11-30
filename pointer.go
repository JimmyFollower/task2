package main

// addTen 接收一个整数指针，在其指向的值上增加 10
func addTen(n *int) {
	if n == nil {
		return
	}
	*n += 10
}
func MultiplyByTwo(s *[]int) {
	for i := range *s {
		(*s)[i] *= 2
	}
}
