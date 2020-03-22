package main

import (
	"fmt"
	"github.com/go-ffmt/ffmt"
)

var (
	pointA = "hello world"
	pointB = "hello world"
	pi *int
	ps *string
)
func main()  {
	ps = &pointA

	fmt.Printf("point is %p\n", ps)
	fmt.Printf("point string is %s\n", *ps)

	testP()

	ffmt.Print(pointA)
	fmt.Printf("point string is %s\n", *ps)
}

func testP () {
	ffmt.Print(pointA)
	{
		pointA := "test"
		ps = &pointA
		ffmt.Print(pointA)
	}
	ffmt.Print(pointA)
}