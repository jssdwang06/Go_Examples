package main

import "fmt"

// 在 Go 中，数组 是一个具有编号且长度固定的元素序列。
func main() {
	// 创建存放 5 个 int 元素的数组 a
	var a [5]int
	// 数组默认值是零值，对于 int 数组来说，元素的零值是 0。
	fmt.Println("emp:", a)

	//  通过 array[index] = value 语法来设置数组指定位置的值，
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("set:", a[4])

	// 内置函数 len 返回数组的长度
	fmt.Println("len:", len(a))

	//在一行内声明并初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 数组类型是一维的，组合构造多维的数据结构。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}
