package main

import "fmt"

func foo() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	my_foo := foo()
	fmt.Println(my_foo())
	fmt.Println(my_foo())
	fmt.Println(my_foo())
}
