package main

import "fmt"

func main() {
	var angka int
	var suara [21]int

	totalMasuk := 0
	suaraSah := 0

	for {
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}

		totalMasuk++

		if angka >= 1 && angka <= 20 {
			suara[angka]++
			suaraSah++
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Println(i, ":", suara[i])
		}
	}
}