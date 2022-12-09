package main

import "fmt"

// 声明全局变量
var g int = 20

func main() {

	fmt.Println("---------------------定义变量---------------------")
	fmt.Printf("g = %d\n", g)
	// 声明局部变量
	var g int = 10
	fmt.Printf("g = %d\n", g)

	fmt.Println("---------------------定义变量---------------------")
	// 形式一
	var n1, n2, n3 string
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	// 形式二
	var name1, age4 = "Tom", 30
	fmt.Println(name1, age4)

	// 形式三
	var age1, age2, age3 int = 31, 32, 33
	fmt.Println(age1, age2, age3)

	// 简短变量
	name2, isBoy, height := "Jay", true, 180.66
	fmt.Println(name2, isBoy, height)

	fmt.Println("------------------接下来是定义常量---------------------")
	// 常量
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	fmt.Println()

	//LENGTH = 14 // 重新赋值会报错：Cannot assign to LENGTH
	//fmt.Printf("LENGTH : %d", LENGTH)

	const a, b, c = 1, false, "str"
	fmt.Println(a, b, c)
}
