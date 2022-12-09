package main

import "fmt"

func main() {
	// --------------------声明数组
	fmt.Println("--------------------声明数组")

	var arr = [3]int{}
	fmt.Println(arr) //[0 0 0]

	var arrStr = [3]string{}
	fmt.Println(arrStr) // [  ]

	// --------------------初始化数组
	arrInt1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrInt1)

	arrStr2 := [1]string{"hello"}
	fmt.Println(arrStr2)

	// 数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
	var balance = [...]float32{1.01, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	//  将索引为 1 和 3 的元素初始化
	balance2 := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance2)

	// ----------------------修改数组
	fmt.Println("----------------------修改数组")

	var arr2 = [2]string{"A", "B"}
	fmt.Println(arr2) //[A B]
	// 修改索引为1的元素值为 C
	arr2[1] = "C"
	fmt.Println(arr2) //[A C]

	// ----------------------访问数组元素
	fmt.Println("----------------------访问数组元素")

	// 遍历数组
	var ins = [...]int{5, 4, 3, 2, 1}
	for i := 0; i < len(ins); i++ {
		fmt.Println(ins[i])
	}

}
