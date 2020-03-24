package main

import "github.com/go-ffmt/ffmt"

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
}