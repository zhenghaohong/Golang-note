package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Tags struct {
	ID int 	`json:"id" gorm:"id"` // 字段 `ID` 默认为主键
	TagName string `json:"tagName" gorm:"tag_name"`
}


// gorm.Model 结构体
/*
 *  基本模型定义 gorm.Model ，包括字段ID、CreatedAt、UpdatedAt、DeletedAt，你可以把它嵌入的模型，或者只写你想要的字段s
*/
type Model struct {
	ID uint `gorm:"primary_key"` // 设置主键
	CreatedAt time.Time  // always update
	UpdatedAt time.Time  // always update
	DeletedAt *time.Time // * first no add
}

// add field `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
type Users struct {
	gorm.Model  // a common struct
	Name string
}

// 表名是结构体名称的复数形式
type User struct {} //默认表名是 `users`

// 设置User的表名为`profiles`
func (Users) TableName() string {
	return "profiles"
}




func main()  {
	db, err := gorm.Open("mysql","root:root@/go_demo1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("connect failure")
	}
	defer db.Close()

	// is exist table:  return bool
	result := db.HasTable("go_role")
	fmt.Println(result)

	// create table 
	db.CreateTable(&Tags{})

	res := db.Model(&Tags{}).AddUniqueIndex("idx_tag_name", "tag_name")

	// var res_value *string
	// res_value = &res
	// fmt.Println(res)
	fmt.Printf("*res_value 变量的值：%v\n", &res)

}