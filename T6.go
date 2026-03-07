package main

import "fmt"

func main() {
	var c1, c2, c3, c4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&c1, &c2, &c3, &c4)

		if !(c1 == "merah" && c2 == "kuning" && c3 == "hijau" && c4 == "ungu") {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}