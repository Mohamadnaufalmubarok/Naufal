package main

import "fmt"

const NMAX = 100000

var data [NMAX]int

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n int, k int) int {
	for i := 0; i < n; i++ {
		if data[i] == k {
			return i
		}
	}

	return -1
}

func main() {
	var n, k int

	// input jumlah data dan angka yang dicari
	fmt.Scan(&n, &k)

	// isi array
	isiArray(n)

	// cari posisi angka
	hasil := posisi(n, k)

	// output
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}