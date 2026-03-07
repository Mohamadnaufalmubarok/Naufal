package main

import (
	"fmt"
)

func main() {
	var beratGram int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratGram)

	kg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaKg := kg * 10000
	tambahan := 0

	if kg >= 10 {
		tambahan = 0
	} else {
		if sisaGram >= 500 {
			tambahan = sisaGram * 5
		} else {
			tambahan = sisaGram * 15
		}
	}

	totalBiaya := biayaKg + tambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, tambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}