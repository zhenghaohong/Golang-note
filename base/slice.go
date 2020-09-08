package main
import "fmt"

/*
*   slice
*   var identifier []type（不需要说明长度)
* 	切片的初始化格式是：var slice1 []type = arr1[start:end] 
* 切片提供了计算容量的函数 cap() 
* 可以测量切片最长可以达到多少：它等于切片的长度 + 数组除切片之外的长度。
* 如果 s 是一个切片，cap(s) 就是从 s[0] 到数组末尾的数组长度。
* 切片的长度永远不会超过它的容量，所以对于 切片 s 来说该不等式永远成立：0 <= len(s) <= cap(s)。
* 
*/
// 切片是一个 长度可变的数组。
func main(){
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
