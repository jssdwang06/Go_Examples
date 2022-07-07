package main

import "fmt"

func main() {
	//单个循环条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// 初始/条件/后续 for 循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// 死循环
	for {
		fmt.Println("loop")
		// 跳出循环
		break
	}
	// continue 进入下次循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
