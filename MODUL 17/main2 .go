package main

import "fmt"

func main() {
	var x string
	var n int

	fmt.Scan(&x)
	fmt.Scan(&n)

	ada := false
	posisi := 0
	jumlah := 0

	for i := 1; i <= n; i++ {
		var data string
		fmt.Scan(&data)

		if data == x {
			if !ada {
				ada = true
				posisi = i
			}
			jumlah++
		}
	}

	fmt.Printf("a. %t\n", ada)
	if ada {
		fmt.Printf("b. %d\n", posisi)
	} else {
		fmt.Printf("b. Tidak ditemukan\n")
	}
	fmt.Printf("c. %d\n", jumlah)
	fmt.Printf("d. %t\n", jumlah >= 2)
}