package main

import (
	"fmt"
	"github.com/go-ffmt/ffmt"
)

var (
	slice [6]int
	map1 map[string]string
)

func main() {
	arr1 := make([]int, 50, 100)
	arr2 := new([10]int)
	arr3 := arr1[50:]
	ffmt.Print(arr1)
	ffmt.Print(arr2)
	ffmt.Print(arr3)

	map2 := new(map[int]int)
	map4 := make(map[int]int)

	map1 = map[string]string{}
	map1["ddd"] = "dddc"
	ffmt.Print(map1)
	map3 := map1

	ffmt.Print(map2)
	ffmt.Print(map3)

	map3["aa"] = "bbb"
	ffmt.Print(map3)
	map4[2] = 2
	ffmt.Print(map4)

	map5 := make(map[int]func() int, 5)

	map5[1] = func() int {
		return 1
	}
	map5[2] = func() int {
		return 2
	}

	ffmt.Print(map5)

	for k, v := range map5 {
		vv := v
		ffmt.Print(k, v, vv)
	}
	//todo 如何将上面函数值 返回出来

	if val, ok := map3["aa"]; ok {
		ffmt.Print(val)
	} else {
		ffmt.Print(222)
	}

	ffmt.Print(map3)
	delete(map3, "aa")
	ffmt.Print(map3)

	//map类型 切片
	map6 := make([]map[int]int, 6)

	for k, _ := range map6 {
		ffmt.Print(k)
		map6[k] = make(map[int]int,2)
		map6[k][2] = 3
	}
	fmt.Printf("The map slice is %v\n", map6)

	map7 := make([]map[int]int, 6)
	fmt.Printf("New map slice is %v\n", map7)

	for _, mv := range map7 {
		mv = make(map[int]int)
		mv[4] = 5
	}
	fmt.Printf("New map slice is %v\n", map7)
}