package main

import (
	"fmt"
	"reflect"
	"github.com/fatih/structs"
)

// UserInfo 用户信息
type UserInfo struct {
	Name string `json:"name" struct:"name"`
	Age  int    `json:"age"  structs:"age"`
	Profile `json:"profile" structs:"profile"`
}

// Profile 配置信息
type Profile struct {
	Hobby string `json:"hobby" structs:"hobby"`
}

func main()  {
	// if_func()
	// for_func()
	// for_mulitpe()
	// range_func()
	u1 := UserInfo{Name: "q1mi", Age: 18, Profile: Profile{"双色球"}}

	//u1 := UserInfo{Name: "q1mi", Age: 18}
	/*
	b, _ := json.Marshal(&u1)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m{
		// 原因 打印出来的 age 类型被自动转成 float64(浮点型 64位)
		// 使用方法 结构体转json  再从 json 转成 map
	      fmt.Printf("key:%v value:%v\n value type:%T\n", k, v, v)
	}
	*/
	/*
		m2, _ := Tomap(&u1,"json")
		for k, v := range m2 {
			fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
		}
	 */

	m := structs.Map(&u1)
	for k, v := range m {
		m["ceshi"] = "测试呀"
		//fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
		fmt.Printf("key:%v value:%v\n", k, v)
	}
}

// Tomap  结构体转为Map[string]interface{} ps: 这个方法的功能 等同于 第三方包 fatih/structs
func Tomap(in interface{}, tagName string) (map[string]interface{}, error){
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {  // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}
	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key; 字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}



// golang range 类似迭代器操作  返回(索引，值)  或(键，值)
// for 循环的 range 格式 可以 对 slice、map、数组、字符串等进行迭代循环。格式如下
// for key, value := range oldMap {
//  	newMap[key] = value
// }
func range_func(){
	//s := "abc"
	//for i := range s {
	//	println(s[i])
	//	//println(i)
	//}
	//
	//for _, c := range s {
	//	println(c)
	//}
	//
	//// 忽略全部返回值  仅迭代
	//for range s {
	//
	//}

	m1 := map[string]string {"a":"5", "type_id":"2"}
	m2 := map[string]string { "ids": "2", "five": "5" }
	fmt.Printf("%#v",m1)
	for k1, v1 := range m1 {
		for k2, v2 := range m2 {
			//if v1 == v2 {
			//	println("有值相同",v2)
			//}
			m1["b2"] = "23"
			println(k1, k2)
			println(v1, v2)
		}
		//println(k1, v1)
	}


}

// for 3种形式 （只有一种使用分号）
func for_func() {
	// 1、for init; condition; post { }
	// 2、for condition {}
	// 3、for {}
	/*
		init 一般为赋值表达式 给控制变量赋初值
		condition 关系表达式或逻辑表达式 循环控制条件
		post 一般为赋值表达式 给控制变量 增量 和 减量
	*/
	s := "abc"
	for i, n:=0, len(s); i < n; i++ {
		println(s[i])
	}

}

// for 嵌套循环
func for_mulitpe() {
	var i, k int
	for i = 2; i < 100; i++ {
		for k=2; k <= (i/k); k++ {
			if i%k==0 {
				break
			}
		}
		if k > (i/k) {
			fmt.Printf("%d  是素数\n", i)
		}
	}


}


// switch
func switch_func()  {

}

// if
func if_func() {
	x := 0
	n:= "abc"
	if x > 0 {
		fmt.Println(n[2])
	}else if x > 1 {
		fmt.Println(n[1])
	}else{
		fmt.Println(n[0])
		fmt.Println(n)
	}
}



