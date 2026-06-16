package main

import (
	"bufio"
	"fmt"
	"os"
)

var teks string
var posisi int

func start() {
	posisi = 0
}

func maju() {
	posisi++
}

func eop() bool {
	return posisi >= len(teks) || teks[posisi] == '.'
}

func cc() byte {
	return teks[posisi]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan teks yang diakhiri titik: ")
	teks, _ = reader.ReadString('\n')

	start()

	jumlahKarakter := 0
	jumlahA := 0
	jumlahLE := 0
	sebelumnya := byte(0)

	for !eop() {
		karakter := cc()

		fmt.Printf("%c", karakter)

		jumlahKarakter++

		if karakter == 'A' {
			jumlahA++
		}

		if sebelumnya == 'L' && karakter == 'E' {
			jumlahLE++
		}

		sebelumnya = karakter
		maju()
	}

	fmt.Println()
	fmt.Println("Jumlah karakter:", jumlahKarakter)
	fmt.Println("Jumlah huruf A:", jumlahA)

	if jumlahKarakter > 0 {
		frekuensiA := float64(jumlahA) / float64(jumlahKarakter)
		fmt.Println("Frekuensi huruf A:", frekuensiA)
	}

	fmt.Println("Jumlah kata LE:", jumlahLE)
}