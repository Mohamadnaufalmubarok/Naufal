package main

import (
	"fmt"
)

// Definisi struct Buku sesuai deskripsi soal
type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

// Menggunakan slice untuk merepresentasikan array dinamis DaftarBuku
type DaftarBuku []Buku

// 1. Procedure untuk menginputkan data buku sebanyak n
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var b Buku
		fmt.Scan(&b.id, &b.judul, &b.penulis, &b.penerbit, &b.eksemplar, &b.tahun, &b.rating)
		*pustaka = append(*pustaka, b)
	}
}

// 2. Procedure untuk mencari dan mencetak buku terfavorit (rating tertinggi) sebelum/sesudah diurutkan
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}
	maxIdx := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	fav := pustaka[maxIdx]
	fmt.Printf("%s %s %s %d\n", fav.judul, fav.penulis, fav.penerbit, fav.tahun)
}

// 3. Procedure mengurutkan buku mengecil (descending) berdasarkan rating dengan INSERTION SORT
func UrutBuku(pustaka DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		// Karena descending, geser elemen yang lebih kecil ke kanan
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

// 4. Procedure mencetak maksimal 5 judul buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Print(pustaka[i].judul)
		if i < limit-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// 5. Procedure mencari buku berdasarkan rating menggunakan BINARY SEARCH
func CariBuku(pustaka DaftarBuku, n int, r int) {
	low := 0
	high := n - 1
	foundIdx := -1

	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			foundIdx = mid
			break // Menemukan salah satu buku
		} else if pustaka[mid].rating < r {
			// Karena terurut menurun (descending), nilai yang lebih besar ada di sebelah kiri
			high = mid - 1
		} else {
			// Nilai yang lebih kecil ada di sebelah kanan
			low = mid + 1
		}
	}

	if foundIdx != -1 {
		b := pustaka[foundIdx]
		fmt.Printf("%s %s %s %d %d %d\n", b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	// Inisialisasi isi pustaka
	pustaka := make(DaftarBuku, 0, n)
	DaftarkanBuku(&pustaka, n)

	// Baris 1 Keluaran: Cetak buku terfavorit
	CetakTerfavorit(pustaka, n)

	// Urutkan array menggunakan Insertion Sort secara Descending
	UrutBuku(pustaka, n)

	// Baris 2 Keluaran: Cetak 5 judul dengan rating tertinggi
	Cetak5Terbaru(pustaka, n)

	// Input rating yang ingin dicari di baris terakhir
	var targetRating int
	fmt.Scan(&targetRating)

	// Baris 3 Keluaran: Cari buku dengan Binary Search
	CariBuku(pustaka, n, targetRating)
}