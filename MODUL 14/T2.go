package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		var jumlah int
		fmt.Scan(&jumlah)

		data := make([]int, jumlah)

		for j := 0; j < jumlah; j++ {
			fmt.Scan(&data[j])
		}

		var ganjil []int
		var genap []int

		// Pisahkan ganjil dan genap
		for _, x := range data {
			if x%2 == 1 {
				ganjil = append(ganjil, x)
			} else {
				genap = append(genap, x)
			}
		}

		// Ganjil urut naik
		sort.Ints(ganjil)

		// Genap urut turun
		sort.Sort(sort.Reverse(sort.IntSlice(genap)))

		// Output
		for _, x := range ganjil {
			fmt.Print(x, " ")
		}

		for _, x := range genap {
			fmt.Print(x, " ")
		}

		fmt.Println()
	}
}