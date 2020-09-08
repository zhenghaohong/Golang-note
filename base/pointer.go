package main
import(
	"fmt"
)
/*
**  指针
* 
*  一个指针变量指向了一个值的内存地址
*  类似于变量和常量  在使用指针
* 
*  指针数组
*
*/
const MAX int = 3

func main() {
	var a int = 6
	var ip *int

	ip = &a // 指针变量的存储地址

	fmt.Printf("a 变量的地址是：%x\n", &a)

	fmt.Printf("ip 变量存储的指针地址：%x\n",ip)

	// 使用指针访问值
	fmt.Printf("*ip 变量的值：%d\n", *ip)

	// 空指针
	/*
		当一个指针被定义后 没有分配到任何变量时，
	*/
	/*
	var ptr *int
	fmt.Printf("prt 的值为：%x\n" ,ptr)
	if ptr == nil {
		fmt.Println("prt 是空指针")
	} else {
		fmt.Println("ptr 不是空指针")
	}
	*/
	ad := []int{10,20,30}
	var i int
	// for i = 0; i < MAX; i++ {
	// 	fmt.Printf("a[%d] = %d\n", i, ad[i] )
	// }

	// 赋值指针数组

	var arr [MAX]*int

	for i=0; i<MAX; i++ {
		arr[i] = &ad[i]
	}

	for i = 0; i<MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *arr[i])
	}




}

