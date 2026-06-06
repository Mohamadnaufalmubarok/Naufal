package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)

	count := 0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.25 {
			count++
		}
	}

	pi := 4.0 * float64(count) / float64(n)

	fmt.Printf("Topping pada Pizza: %d\n", count)
	fmt.Printf("PI : %.10f\n", pi)
}