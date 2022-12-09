package main

import "fmt"

func main() {

	var a int = 20

	// 指针变量的存储地址
	var ip *int = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)
	// 使用指针访问值
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var intArr [3]int // int占8个字节
	//定义完数组后，数组的各个元素就已经有默认值了
	fmt.Printf("intArr的地址%p\n", &intArr)
	fmt.Printf("intArr[0]的地址是%p\n", &intArr[0])
	fmt.Printf("intArr[1]的地址是%p\n", &intArr[1])
	fmt.Printf("intArr[2]的地址是%p\n", &intArr[2])
}
