package main

import "fmt"

func main() {
	//使用内建函数 make 创建一个长度不为 0 的空 slice
	s := make([]string, 3) // 长度为3， 初始值为零

	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("set:", s[2])

	// len 返回 slice 的长度
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片
	l := s[2:5]
	fmt.Println("sl1:", l)

	l1 := s[:5]
	fmt.Println("sl2:", l1)

	l2 := s[2:]
	fmt.Println("sl3:", l2)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice 可以组成多维数据结构。内部的 slice 长度可以不一致，这一点和多维数组不同
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
