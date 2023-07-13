package main

import (
	"fmt"
)

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("result : ", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}
}
func loopDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("i:", i)
	}
}
func main() {
	// fmt.Println("Welcom to cal")
	// defer fmt.Println("End")
	// defer add(20, 30)
	// defer add(15, 30)
	// add(50, 30)
	loop()
	loopDefer()
}
