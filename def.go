package main

import "fmt"

func main() {
	fmt.Println("Todos os números que são divisíveis por 3, entre 1 a 100, estão logo abaixo:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}

	}
}
