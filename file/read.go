package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("readme.txt")
	if err != nil {
		fmt.Println("err")
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	fmt.Println("scanner")
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d : %s\n", count, line)
		count++
	}
}
