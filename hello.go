package main

import "fmt"

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func main() {
	trace("i")
	defer untrace("i")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
}
