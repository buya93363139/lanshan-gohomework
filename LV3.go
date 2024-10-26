package main

import (
	"fmt"
)

func main() {

	for i := 2; i <= 100; i++ {

		var flag = true

		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}

}
