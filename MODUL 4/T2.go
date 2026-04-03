package main

import (
	"fmt"
)

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, namaPemenang string
	var soal, skor int
	var maxSoal int = -1
	var minSkor int = 9999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}
		
		hitungSkor(&soal, &skor)

		if soal > maxSoal {
			maxSoal = soal
			minSkor = skor
			namaPemenang = nama
		} else if soal == maxSoal && skor < minSkor {
			minSkor = skor
			namaPemenang = nama
		}
	}

	if maxSoal != -1 {
		fmt.Println(namaPemenang, maxSoal, minSkor)
	}
}