package hello

import "fmt"

func Hello() {
	fmt.Println(hello(), morning())
}

func hello() string {

	return "Hello there!"
}

func morning() string {

	return "Good morning!"
}