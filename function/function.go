package main

import "fmt"

func hello() {
	fmt.Println("Hello Borntodev")
}

func plus(value1 int, value2 int) int {
	return value1 + value2
}

func plus3(value1 int, value2 int, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()
	result := plus(1, 5)
	fmt.Println("result:", result)

	result2 := plus3(1, 5, 3)
	fmt.Println("result2:", result2)
}
