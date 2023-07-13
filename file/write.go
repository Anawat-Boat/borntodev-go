package main

import "os"

func main() {
	deta1 := []byte("hello\n borntodev")
	err := os.WriteFile("data.txt", deta1, 0644)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("employeeName")

	defer f.Close()
	data2 := []byte("anawat\n monprachak")
	os.WriteFile("employeeName.txt", data2, 0644)
}
