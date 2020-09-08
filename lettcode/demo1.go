package main

// 10 classic calculation for golang

import (
    "fmt"
)

func main() {
    arr1 := []int{20, 2, 7, 11, 15}
    // fmt.Println(twoSum(arr1,26))

    // fmt.Println(bubbleSort(arr1))

    // #quickSort#
    // fmt.Println(quickSort(arr1)) // method one
    quickSort2(arr1)                // method two
    fmt.Println(arr1)

}

// 给定一个整数数组 nums 和一个目标值 target，
// 请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
func twoSum(nums []int, target int) []int {
    l := len(nums)
    for i:=0;i<l;i++{
        for j:=i+1;j<l;j++{
            if nums[i]+nums[j] == target{
                return []int{i,j}
            }
        }
    }
    return []int{}
}

// bubbleSort
// step 1 : 比较相邻的元素。如果第一个比第二个大，就交换他两
// step 2 : 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// step 3 : 针对所有的元素重复以上的步骤，除了最后一个。
// step 4 : 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

func bubbleSort (arr []int ) []int {
    length := len(arr)
    for i := 0; i < length; i++ {
        for j := 0; j < length-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr

}



// quickSort1
// step 1: 从数列中挑出一个元素，称为 “基准”（pivot）;
// step 2: 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
//         在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
// step 3: 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；


func quickSort(arr []int) []int {
    return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
    if left < right {
        partitionIndex := partition(arr, left, right)
        _quickSort(arr, left, partitionIndex-1)
        _quickSort(arr, partitionIndex+1, right)
    }
    return arr
}

func partition (arr []int, left, right int) int {
    pivot := left
    index := pivot + 1
    for i := index; i <= right; i++ {
        if arr[i] < arr[pivot] {
            swap(arr, i, index)
            index += 1
        }
    }
    swap(arr, pivot, index-1)
    return index - 1
}

func swap(arr []int, i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}

// quickSort 2 for liwenzhou

func quickSort2(data []int) {
    if len(data) <= 1 {
        return
    }
    base := data[0]
    l, r := 0, len(data)-1
    for i := 1; i <= r; {
        if data[i] > base {
            data[i], data[r] = data[r], data[i]
            r--
        } else {
            data[i], data[l] = data[l], data[i]
            l++
            i++
        }
    }
    quickSort(data[:l])
    quickSort(data[l+1:])
}


// insertionSort
// introduction
//     插入排序和冒泡排序一样，也有一种优化算法，叫做拆半插入。
// step
// step 1: 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
// step 2: 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
//        （如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）

func insertionSort(arr []int) []int {
    for i := range arr {
        preIndex := i - 1
        current := arr[i]
        for preIndex >= 0 && arr[preIndex] > current {
            arr[preIndex+1] = arr[preIndex]
            preIndex -= 1
        }
        arr[preIndex+1] = current
    }
    return arr
}











