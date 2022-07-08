package main

import "fmt"

func main() {
	// range 用来迭代各种数据结构

	nums := []int{2, 3, 4}
	sum := 0

	// 忽略索引
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 索引
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 迭代键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n\n", k, v)
	}

	// 遍历 map 的键
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range 在字符串中迭代 unicode 码点(code point)。
	//第一个返回值是字符的起始字节位置，然后第二个是字符本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
