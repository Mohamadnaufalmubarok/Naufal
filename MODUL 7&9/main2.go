package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("a. Semua elemen: ")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Print("d. Indeks kelipatan ", x, ": ")
	for i := 0; i < len(arr); i++ {
		if x != 0 && i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var delIdx int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&delIdx)
	arr = append(arr[:delIdx], arr[delIdx+1:]...)
	fmt.Print("e. Setelah dihapus: ")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += float64(arr[i])
	}
	avg := sum / float64(len(arr))
	fmt.Printf("f. Rata-rata: %.2f\n", avg)

	var sumSq float64
	for i := 0; i < len(arr); i++ {
		diff := float64(arr[i]) - avg
		sumSq += diff * diff
	}
	stdDev := math.Sqrt(sumSq / float64(len(arr)))
	fmt.Printf("g. Standar Deviasi: %.2f\n", stdDev)

	var target int
	fmt.Print("Masukkan bilangan untuk dicari frekuensinya: ")
	fmt.Scan(&target)
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			count++
		}
	}
	fmt.Printf("h. Frekuensi %d muncul sebanyak: %d kali\n", target, count)
}