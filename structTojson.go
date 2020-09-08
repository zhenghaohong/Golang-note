package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int64
}
/*
	来自：李文周博客 https://www.liwenzhou.com/posts/Go/json_tricks_in_go/
	json.Marshal() 		系列化
	json.Unmarshal() 	反序列化
	struct 结构体
 */
func main()  {

	p1 := Person{
		Name: "郑特帅",
		Age:  18,
	}
	// struct 转 json string
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("struct 转 json 失败 json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("%s\n",b)

	// json string 转 struct
	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json 转 struct 失败 json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)

	// 结构体的嵌套 事例1
	nestedStructDemo()
}

// omitempty 表示 值 为空 时  不返回该字段

type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
	Profile
}


type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
	Test 	int    `json:"test"`
}

// struct 嵌套
func nestedStructDemo() {
	u1 := User{
		Name:  "七米",
		// Email: "",
		Hobby: []string{"足球", "双色球"},
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}



