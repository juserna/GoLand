package main

import "fmt"

func main() {
	// closure
	// anonymous functions

	func(i int) int {
		return i + 10
	}(8)
}

/*
reserved word "func"
name
arguments (optional)
return type (optional in case of nothing)
multiple return type (int, string, error)
named return type (use it with naked return)*/

func f1() {
	fmt.Println("No args and no return")
}

func f2(i int) int {
	return i + 10
}

func f3(i, j int) (int, int) {
	return i + j, i * j
}

// naked return
func f4(i, j int) (sum int, multi int) {
	sum = i + j
	multi = i * j
	return // naked return
	// return sum, multi (opcion valida)
}

func invariantArgs(i ...int) (sum int) {
	for _, v := range i {
		sum += v
	}
	return
}
