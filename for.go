package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("num %d\n", i)
	}


	var i = 10
	for i > 5 {
		i = i -1
		fmt.Printf("this is %d\n", i)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}

	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

	s := ""
	for ; s != "aaaaa"; {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}


	str := "test 中文"

	for k,v := range str {
		fmt.Printf("the key is %d, and the value is %c\n",k, v)
	}

	for index, rune := range str {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}

