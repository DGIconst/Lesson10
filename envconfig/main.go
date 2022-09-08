package main

import (
	"confpkg/factorial"
	"confpkg/hello"
	"fmt"
)

func main() {
	fmt.Println(factorial.Factorial(6))
	fmt.Println(factorial.IterativeFactorial(6))
	hello.Hello()
}
