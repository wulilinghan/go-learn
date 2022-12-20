package main

import "fmt"

func main() {
	// 定义 map ，这个时候为nil 也不能进行添加元素，获取元素等操作
	var letterMap map[string]string
	fmt.Println(letterMap)
	if letterMap == nil {
		fmt.Println("letterMap is nil")
	}
	// 初始化
	letterMap = make(map[string]string)
	letterMap["A"] = "a"
	letterMap["B"] = "b"
	letterMap["C"] = "c"

	// 遍历
	for key := range letterMap {
		fmt.Println(key, " - ", letterMap[key])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := letterMap["American"]
	fmt.Println(capital)
	fmt.Println(ok)
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	// 删除元素
	delete(letterMap, "A")
	fmt.Println(letterMap)

}
