package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

/*
	可以方便地从一个 JSON 串中读取值。同时它也支持各种查询、汇总统计等功能。今天我们再介绍一个类似的库gjson。
	在上一篇文章Go 每日一库之 buntdb中我们介绍过 JSON 索引，内部实现其实就是使用gjson这个库。
	gjson实际上是get + json的缩写，用于读取 JSON 串，同样的还有一个sjson（set + json）库用来设置 JSON 串
*/
const jsonStr = `{
  "name":{"first":"Tom", "last": "Anderson"},
  "age": 37,
  "children": ["Sara", "Alex", "Jack"],
  "fav.movie": "Dear Hunter",
  "friends": [
    {"first": "Dale", "last":"Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`


func main() {
	jsonstring := `{"data":{"name":"zhengteshuai","sex":"man"},"code":"200","msg":"power"}`
	//fmt.Println("read json data ")
	//json := `{"name":{"first":"www.topgoer.com","last":"dj"},"age":18}`
	theMap := gjson.Get(jsonstring, "data")
	fmt.Println("the map", theMap)
	theName := gjson.Get(jsonstring, "data.name")
	fmt.Println("the name:", theName.String())
	age := gjson.Get(jsonstring, "code")
	fmt.Println("code:", age.Int())

}


