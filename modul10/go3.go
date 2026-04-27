package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(data arrBalita, n int, min *float64, max *float64) {
	*min = data[0]
	*max = data[0]

	for i := 1; i < n; i++ {
		if data[i] < *min {
			*min = data[i]
		}
		if data[i] > *max {
			*max = data[i]
		}
	}
}

func rerata(data arrBalita, n int) float64 {
	var jumlah float64

	for i := 0; i < n; i++ {
		jumlah = jumlah + data[i]
	}

	return jumlah / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(data, n))
}