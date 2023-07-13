package main

import "fmt"

var productName [4]string
var productPrice [4]float32

func main() {
	//Array
	productName[0] = "MacBook"
	productName[1] = "IPhone"
	productName[2] = "IMac"
	productName[3] = "IPad"

	productPrice := [4]float32{4000, 3000, 20000, 2000}
	fmt.Println(productName)
	fmt.Println(productPrice)
}
