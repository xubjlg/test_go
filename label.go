package main

import "fmt"

func main() {
//LABLE1:
//	for i := 1; i < 2; i++ {
//		for j := 1; j < 10; j++ {
//			if j > 4 {
//				goto LABLE1
//			}
//			ffmt.Print(j)
//		}
//	}


	a := 1
	b := 3
	//LAB :
		a =+ b
		fmt.Printf("a is %d and b is %d", a, b)
		a += b
		fmt.Printf("a is %d and b is %d", a, b)
	//goto LAB

	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 { break }
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")

	for i := 0; i<7 ; i++ {
		if i%2 == 0 { continue }
		fmt.Println("Odd:", i)
	}
}