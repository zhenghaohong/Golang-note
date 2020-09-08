package main

import (
	"fmt"
)
func main()  {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 3
	}
	// 多维数组不能直接打印 需要遍历打印

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("index  %d  is  %d\n",i, arr1[i])
	}
	
	// var 数组名称  [数组数量]数组类型
	
}

