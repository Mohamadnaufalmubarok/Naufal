package main

import (
	"fmt"
)

func main() {
	var b int
	var jumlahFaktor int

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	if b <= 1 {
		fmt.Println("Masukkan bilangan bulat lebih dari 1.")
		return
	}

	fmt.Print("Faktor: ")

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
			jumlahFaktor++ 
		}
	}
	fmt.Println() 
	var isPrima bool
	if jumlahFaktor == 2 {
		isPrima = true
	} else {
		isPrima = false
	}
	fmt.Printf("Prima: %t\n", isPrima)
}