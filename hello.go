package main

import (
	"fmt"
)

func main() {
	var a int = 1
	var b int = 1
	fmt.Println(a)
	fmt.Println(b)

	for i := 0; i < 6; i++ {
		fmt.Println(b)
		tmp := a
		a = b
		b = tmp + a

	}

}
