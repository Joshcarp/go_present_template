package main

import "fmt"

func main() {

	for i := 0; i < 16777217; i++ {
		if int(float64(i)) != i {
			fmt.Println(i, float32(i))
			panic("")
		}
	}
	fmt.Println("Finished")
}
