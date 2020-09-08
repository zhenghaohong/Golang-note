package main

import (
	"fmt"
	"strings"
	"time"
)



func main()  {


	//  0801  8  2
	timeStr := "2020-08-22"
	a := strings.Split(timeStr, "-")
	fmt.Printf("%d****%v****%v\n", len(a), a[0],a[1])


	// 8 和 2 
	ur := fmt.Sprintf("%s  %s",strings.Replace(a[1], "0", "", -1 ),strings.Replace(a[2], "0", "", -1 ))

	// fmt.Printf("ur type:%T\n", ur)
	fmt.Println(ur)


	month:=time.Now().Month()//time.Now().Month().String()
	day:=time.Now().Day()
	month+=1
	
	if month == 9 {
		month =14
	}
	fmt.Println(month)
	fmt.Printf("%v-%v\n",int(month),day)
	// fmt.Println(int(month),"-",day)
	// s := "123lafaldsjglad123lkfasdf123djfal123lkdjga123lksjfla123l"
    // old := "123"
    // new := "888"

 	// // n < 0 ,用 new 替换所有匹配上的 old；n=-1:  888lafaldsjglad888lkfasdf888djfal888lkdjga888lksjfla888l
	// fmt.Println("n=-1: ", strings.Replace(s, old, new, -1 ))
	// fmt.Println("result: ", strings.Replace(ur, "0", "", -1 ))


}
