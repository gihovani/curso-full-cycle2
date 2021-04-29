package main

import "fmt"

func main() {
	fmt.Println(Sum(10, 10))
}

func Sum(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Div(a int, b int) int {
	return a / b
}

func Times(a int, b int) int {
	return a * b
}

func Pow(a int, b int) int {
    ret := a
    for i := 1; i < b; i++ {
        ret *= a
    }
	return ret
}
