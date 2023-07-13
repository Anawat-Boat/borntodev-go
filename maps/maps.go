package main

import "fmt"

var product = make(map[string]int)

func main() {
	fmt.Println("product:", product)

	//add
	product["Macbook"] = 40000
	product["IPhone"] = 30000
	product["IPad"] = 25000
	fmt.Println("product:", product)

	// delete
	delete(product, "IPad")
	fmt.Println("product:", product)

	// change
	product["Macbook"] = 20000
	fmt.Println("product:", product)

	// get value
	value := product["Macbook"]
	fmt.Println("value:", value)

	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println("courseName:", courseName)
}
