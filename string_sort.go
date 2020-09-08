package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {

	// define a array
	arrList := []string{"3","5","1","9","6","7"}
	//fmt.Printf("Before: %v\n", arrList)
	fmt.Println("before \n",arrList)

	sort.Slice(arrList, func(i, j int) bool {
		numA, _ := strconv.Atoi(arrList[i])
		numB, _ := strconv.Atoi(arrList[j])
		return numA < numB
	})

	fmt.Println("after \n",arrList)

}
