package main

import "fmt"

func main() {
	var x, y, i, j, index int
	var ikan [1000]float64
	var total, jumlah float64

	fmt.Print("Masukkan jumlah wadah: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah ikan per wadah: ")
	fmt.Scan(&y)

	for i = 0; i < x*y; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	index = 0

	for i = 0; i < x; i++ {
		total = 0

		for j = 0; j < y; j++ {
			total += ikan[index]
			jumlah += ikan[index]
			index++
		}

		fmt.Printf("Total berat wadah %d: %.2f\n", i+1, total)
	}

	fmt.Printf("Rata-rata berat ikan: %.2f\n", jumlah/float64(x*y))
}