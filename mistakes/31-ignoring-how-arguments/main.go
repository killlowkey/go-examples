package main

import "fmt"

// main 以下例子来源于 100 Go Mistakes and How to Avoid Them
// #31: Ignoring how arguments are evaluated in range loops
func main() {

}

func rangeExample() {
	s := []int{0, 1, 2}

	// 这里是使用 s 的副本来进行遍历，所以向 s 追加数据不会影响遍历次数
	for range s {
		fmt.Println("hello")
		s = append(s, 10)
	}

	fmt.Println("s: ", s)
}

func indexExample() {
	s := []int{0, 1, 2}
	// 使用本身的 s 进行遍历，向s追加数据，也会导致遍历一直死循环
	for i := 0; i < len(s); i++ {
		s = append(s, 10)
	}
}
