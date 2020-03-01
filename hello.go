package main

import (
	"fmt"
	"github.com/go-ffmt/ffmt"
)

var (
	a = 3
	b int
)

func main() {

	//main1()

	//测试 作用域 范围
	//main2()


	//测试 全局变量 
	main3()
	ffmt.Print(b)


}


func main1() {
	a := [10]int64{1, 2, 3, 4, 5, 6, 7, 8}
	//b := map[]{}
	fmt.Println("Hello world")

	fmt.Printf("%s\n",a)

	ffmt.Print(a)
}

func main2() {
	ffmt.Print(a)
	a := 1

	ffmt.Print(a)

	{
		a := 2
		ffmt.Print(a)
	}

	ffmt.Print(a)
}

func main3() {
	ffmt.Print(b)
	b = 3
	{
		b := 4
		ffmt.Print(b)
	}
	ffmt.Print(b)
}