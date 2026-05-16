package main

import "fmt"

func main() {
	var angka int
	var suara [21]int

	totalMasuk := 0
	suaraSah := 0

	// input data
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

	// mencari ketua dan wakil
	ketua := 1
	wakil := 1

	for i := 1; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i
		} else if i != ketua && suara[i] > suara[wakil] {
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}