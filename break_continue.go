package main

import "github.com/go-ffmt/ffmt"

func main()  {
	for i :=1;i <= 4; i++ {
		for j := 1; j<= 10 ;j++ {
			if j > 6 {
				break
			}
			ffmt.Print(j)
		}
	}
}
