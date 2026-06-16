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

func main() {
	domino := buatDominoes()

	kocokKartu(&domino)

	rangkaian := []Domino{ambilKartu(&domino)}

	fmt.Println("Kartu awal:", rangkaian[0].kiri, "|", rangkaian[0].kanan)

	ujung := rangkaian[0].kanan

	for domino.sisa > 0 {
		kartu := ambilKartu(&domino)

		if kartu.kiri == ujung || kartu.kanan == ujung {
			rangkaian = append(rangkaian, kartu)

			if kartu.kiri == ujung {
				ujung = kartu.kanan
			} else {
				ujung = kartu.kiri
			}

			fmt.Println("Ditambahkan:", kartu.kiri, "|", kartu.kanan)
		}
	}

	fmt.Println("Jumlah kartu dalam rangkaian:", len(rangkaian))
}