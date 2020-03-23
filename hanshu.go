package main

import (
	"fmt"
	"github.com/go-ffmt/ffmt"
	"strconv"
)

func main() {
	fmt.Printf(" the number is 2 * 3 * 5 = %d\n", Multiply(2, 3, 5))

	num1, num2 := getV(2, 3)
	fmt.Printf("this num1 : %d and num2 %d\n", num1, num2)

	num1, num2 = getV1(2, 3)
	fmt.Printf("this num1 : %d\n and num2 %d\n", num1, num2)

	num1, num2 = getV2(2, 3)
	fmt.Printf("this num1 : %d\n and num2 %d\n", num1, num2)

	var st *string
	av3 := "a"
	st = &av3
	getV3P(2, 4, st)
	ffmt.Print(st)


	var st1 * int
	n := 0
	st1 = &n
	getV4P(4,5, st1)
	fmt.Printf("num is %d\n", *st1)

	var st2 * int
	n1 := 0
	st2 = &n1
	getV5P(4,5, st2)
	fmt.Printf("num is %d\n", *st2)

	var st3 * int
	n2 := 0
	st3 = &n2
	getV6P(4,5, st3)
	fmt.Printf("num is %d\n", *st3)

	moreArg("aa","bbb","ccc", "ddd")

	defer func() {
		ffmt.Print("defer start")
		if error := recover(); error != nil {
			ffmt.Print("get error")
		}
		ffmt.Print("defer end")
	}()

	def()

	diGui(10)

}

func Multiply(a,b,c int) int {
	return a * b * c
}



//非命名返回值
func getV(a,b int) (int, int) {
return a + 1,b +1
}

//命名返回值
func getV1(a int, b int) (c int, d int) {
	c = a * a
	d = b + a
	return d, c
}
//命名返回值
func getV2(a int, b int) (c int, d int) {
	c = a * a
	d = b + a
	return
}

func getV3P(a, b int, c *string)  {
	d := strconv.Itoa(a * b)
	*c = d
}

func getV4P(a, b int, c *int)  {
	d := a * b
	c = &d
}
func getV5P(a, b int, c *int)  {
	d := a * b
	*c = d
}

func getV6P(a,b int, c *int) int {
	d := a * b
	c = &d
	return *c
}

func moreArg(a ...string)  {
	for _,v := range a {
		fmt.Printf("print value : %s\n",v)
	}
}

func def()  {
	ffmt.Print("start")
	//panic("error 555")
	ffmt.Print("end")
}

//递归
func diGui(i int)  {

	for j := i; j <= i+10; j++ {
		ffmt.Print(j)
		res := dg(j)
		fmt.Printf("di gui ：%d\n", res)
	}
}

func dg(i int) (res int){
	if i <= 1 {
		res = i
	} else {
		res = dg(i -1) + dg(i - 2)
	}
	return
}

