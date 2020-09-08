package main

import (
	"os"
	"fmt"
	)

func main(){

	err:=os.Mkdir("./dir1",os.ModePerm)
	if err!=nil{
		fmt.Println(err)
	}




}
