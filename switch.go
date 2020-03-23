package main

import (
	"fmt"
	"strconv"
)

func main() {

	for a := 1;a <= 10;a++ {
		fmt.Printf(ss(&a))
		fmt.Print("\n")
	}
}

func ss(i *int) string{
	switch *i {
	case 1:
		return "test 1"
	case 2:
		return strconv.Itoa(sss(i))
	default:
		return "test string" + strconv.Itoa(*i)
	}
}

func sss(i *int)  int{
	return *i + 10
}