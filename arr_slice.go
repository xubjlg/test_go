package main

import (
	"fmt"
	"github.com/go-ffmt/ffmt"
)

var (
	arr  [10]int
	arr1 [5]string
)

func main()  {
	arr[1] = 6

	arr1[2] = "test2"
	arr1[4] = "2"

	ffmt.Print(arr)
	ffmt.Print(arr1)

	var2 := [...]string{"a", "b", "c"}

	for _,i := range var2 {
		ffmt.Print(i)
	}

	var arr2 = new([5]int)

	var arr3  [5]int

	for i := range arr2 {
		ffmt.Print(i)
	}

	arr3 = *arr2

	arr3[2] = 100

	_, _ = ffmt.Print(arr3)
	_, _ = ffmt.Print(*arr2)

	var arr4 = [10]int{1,2,3}
	for k,v := range arr4 {
		fmt.Printf("key is %d, value is %d\n", k, v)
	}

	ffmt.Print(len(arr4))

	arr5 := []int{1,2,3,4,5}
	arr6 := arr5[:]
	ffmt.Print(arr6)

	//arr7 := arr6[:3] + arr6[3:]
	//ffmt.Print(arr7)


	//make 初始化  new 函数分配内存
	//var arr8 = make([]int, 50,100)
	//var arr9 = new([100]int)[:50]

	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		ffmt.Print(i)
		arr1[i] = i
	}

	// print the slice
	fmt.Print(len(slice1))
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


	var arr10 = [10]int{0,1,2,3,4,5,6,7,8,9}
	arr11 := arr10[4:7]
	ffmt.Print(arr11)
	arr11 = arr11[0:6]
	ffmt.Print(arr11)

	ffmt.Print(len(arr11[5:5]))
	ffmt.Print(len(arr11[5:6]))


	var arr12 = []int{1,2,3}
	arr12 = append(arr12, 4,5,6)
	ffmt.Print(arr12)
}

