package main

import "fmt"

func main() {

	courseName := []string{"Java", "Python"}
	fmt.Println(courseName)
	//append add
	courseName = append(courseName, "C")
	fmt.Println(courseName)
	courseName = append(courseName, "C#", "HTML")
	fmt.Println(courseName)
	//
	courseWeb := courseName[4:5]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
