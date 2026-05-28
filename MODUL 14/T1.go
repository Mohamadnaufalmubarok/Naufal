package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan algoritma Selection Sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Asumsikan elemen saat ini adalah yang terkecil
		minIdx := i
		
		// Cari elemen yang lebih kecil di sisa array
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		
		// Tukar elemen terkecil yang ditemukan dengan elemen di posisi i
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	
	// Membaca input n (banyaknya daerah)
	fmt.Scan(&n)

	// Looping sebanyak n daerah
	for i := 0; i < n; i++ {
		var m int
		
		// Membaca m (banyaknya rumah kerabat di daerah tersebut)
		fmt.Scan(&m)

		// Membuat slice untuk menyimpan nomor rumah sebanyak m
		rumah := make([]int, m)
		
		// Membaca nomor-nomor rumah
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		// Mengurutkan nomor rumah menggunakan Selection Sort
		selectionSort(rumah)

		// Mencetak hasil yang sudah diurutkan
		for j := 0; j < m; j++ {
			if j > 0 {
				fmt.Print(" ") // Tambahkan spasi antar angka
			}
			fmt.Print(rumah[j])
		}
		fmt.Println() 
	}
}