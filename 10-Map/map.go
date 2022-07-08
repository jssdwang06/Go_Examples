package main

import "fmt"

// python 中的字典
func main() {

	// 使用make函数创建一个空map make(map[key-type]val-type)
	m := make(map[string]int)

	// 设置键值对 name[key] = val
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// 使用 name[key]来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// 返回键值对数量
	fmt.Println("len:", len(m))

	// 内建函数 delete 删除 map 中的键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 使用 _ 忽略不需要的值
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 声明并初始化一个新 map
	n := map[string]int{"foo": 1, "bar": 2}

	// 以 map[k:v k:v] 的格式输出
	fmt.Println("map:", n)
}
