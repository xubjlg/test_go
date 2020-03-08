package main

import (
	"fmt"
	"github.com/go-ffmt/ffmt"
	"time"
)

var (
	a = 3
	b int
)

func main() {
	//print
	main1()

	//测试 作用域 范围
	main2()

	//测试 全局变量
	main3()
	_, _ = ffmt.Print(b)

	b = 5
	_, _ = ffmt.Print(b)

	baseTime()
}


func main1() {
	a := [10]int64{1, 2, 3, 4, 5, 6, 7, 8}
	//b := map[]{}
	fmt.Println("Hello world")

	fmt.Printf("%s\n",a)

	_, _ = ffmt.Print(a)
}

func main2() {
	_, _ = ffmt.Print(a)
	a := 1

	_, _ = ffmt.Print(a)
	{
		a := 2
		_, _ = ffmt.Print(a)
	}

	_, _ = ffmt.Print(a)
}

func main3() {
	_, _ = ffmt.Print(b)

	b = 3
	{
		b := 4
		_, _ = ffmt.Print(b)
	}

	_, _ = ffmt.Print(b)
}

func baseTime() {
	t := time.Now()

	_, _ = ffmt.Print(t)
//	fmt.Printf(t)

	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(),t.Month(),t.Year())

	t = time.Now().UTC()
	fmt.Println(t)

	tt := 5 * time.Second

	_, _ = ffmt.Print(tt)

	week := 60 * 60 * 24 * 7 * 1e9

	newWeek := t.Add(time.Duration(week))
	fmt.Println(newWeek)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println(time.Now())

	lo, _ := time.LoadLocation("")
	fmt.Println(lo)


	ticker  := time.Tick(1 * time.Second)

	for _ = range ticker {
		fmt.Println(1 )
	}


	newTicker := time.NewTicker(5 * time.Second)

	newTicker.Stop()  // todo  如何开启 如何停用


newTime(1 * time.Second)

}


func newTime(d time.Duration) time.Timer {
	t := time.Timer{}
	fmt.Println("time")

	return t
}