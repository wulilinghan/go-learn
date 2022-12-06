package main

import "fmt"

func main() {

	var str string
	var i int
	var boolVal bool
	var f float32

	fmt.Println("str: " + str)
	fmt.Println(i)
	fmt.Println(boolVal)
	fmt.Println(f)

	fmt.Println("-----------------------------------------------")

	str = "Y"
	i = 100
	boolVal = true
	fmt.Println("str: " + str)
	fmt.Println(i)
	fmt.Println(boolVal)

	fmt.Println("-----------------------------------------------")

	fmt.Println(&str) //打印地址值
	var strCopy string
	strCopy = str
	fmt.Println(&str)
	fmt.Println(&strCopy)

}
