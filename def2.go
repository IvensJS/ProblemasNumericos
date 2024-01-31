package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pin, Pan", i)
		}
		if i%3 == 0 {
			fmt.Println("Pin", i)
		}
		if i%5 == 0 {
			fmt.Println("Pan", i)
		}
	}
}

// A escrita pin e pan, estão aparecendo separadamente nos números que ambos são divisores, preciso arrumar isso.