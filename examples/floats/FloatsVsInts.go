package main

import "fmt"

func main() {
	for i := 0; i < 33554490; i++ {
		if int(float32(i)) != i {
			fmt.Println(i, float32(i))
			panic("")
		}
	}
	fmt.Println("Finished")
}
