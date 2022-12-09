package main

import "fmt"

func main() {

	var slice1 = make([]int, 3, 5)
	fmt.Printf("slice1 变量的地址是: %p\n", &slice1)

	fmt.Println("len(slice1)=", len(slice1))
	fmt.Println("cap(slice1)=", cap(slice1))
	fmt.Println("slice1 = ", slice1)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3

	fmt.Println("重新赋值的slice1 = ", slice1)
	fmt.Printf("slice1变量的地址是: %p\n", &slice1)

	// 将 slice1 中从下标 0 到 1 下的元素创建为一个新的切片
	var ins = slice1[0:2]
	fmt.Println("ins = ", ins)
	fmt.Printf("ins 变量的地址是: %p\n", &ins)

	// 将 slice1 中从下标 2 到结尾最后元素创建为一个新的切片
	var ins2 = slice1[2:]
	fmt.Println("ins2 = ", ins2)
	fmt.Printf("ins2 变量的地址是: %p\n", &ins2)

	// 将 slice1 从第一个元素到下标为1的元素创建为一个新的切片。
	var ins3 = slice1[:2]
	fmt.Println("ins3 = ", ins3)
	fmt.Printf("ins3 变量的地址是: %p\n", &ins3)
}
