package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	kiri  int
	kanan int
	balak bool
}

type Dominoes struct {
	kartu []Domino
	sisa  int
}

func buatDominoes() Dominoes {
	var d Dominoes

	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			d.kartu = append(d.kartu, Domino{
				kiri:  i,
				kanan: j,
				balak: i == j,
			})
		}
	}

	d.sisa = len(d.kartu)
	return d
}

func kocokKartu(d *Dominoes) {
	rand.Seed(time.Now().UnixNano())

	for i := len(d.kartu) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.kartu[i], d.kartu[j] = d.kartu[j], d.kartu[i]
	}
}

func ambilKartu(d *Dominoes) Domino {
	kartu := d.kartu[0]
	d.kartu = d.kartu[1:]
	d.sisa--

	return kartu
}

func gambarKartu(d Domino, suit int) int {
	if suit == 1 {
		return d.kiri
	}

	return d.kanan
}

func nilaiKartu(d Domino) int {
	return d.kiri + d.kanan
}

func galiKartu(d *Dominoes, kartuAwal Domino) Domino {
	for d.sisa > 0 {
		kartu := ambilKartu(d)

		if kartu.kiri == kartuAwal.kiri ||
			kartu.kanan == kartuAwal.kiri ||
			kartu.kiri == kartuAwal.kanan ||
			kartu.kanan == kartuAwal.kanan {
			return kartu
		}
	}

	return Domino{}
}

func sepasangKartu(kartu1 Domino, kartu2 Domino) bool {
	return nilaiKartu(kartu1)+nilaiKartu(kartu2) == 12
}

func main() {
	domino := buatDominoes()

	fmt.Println("Jumlah kartu awal:", domino.sisa)

	kocokKartu(&domino)

	kartu1 := ambilKartu(&domino)
	fmt.Println("Kartu pertama:", kartu1.kiri, "|", kartu1.kanan)

	kartu2 := galiKartu(&domino, kartu1)
	fmt.Println("Kartu hasil gali:", kartu2.kiri, "|", kartu2.kanan)

	fmt.Println("Sisi kiri kartu pertama:", gambarKartu(kartu1, 1))
	fmt.Println("Sisi kanan kartu pertama:", gambarKartu(kartu1, 2))
	fmt.Println("Nilai kartu pertama:", nilaiKartu(kartu1))

	if sepasangKartu(kartu1, kartu2) {
		fmt.Println("Total nilai kedua kartu adalah 12")
	} else {
		fmt.Println("Total nilai kedua kartu bukan 12")
	}

	fmt.Println("Jumlah kartu tersisa:", domino.sisa)
}