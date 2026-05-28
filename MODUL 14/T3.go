package main

import "fmt"

func main() {
	var data []int
	var x int

	// Input sampai bilangan negatif
	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		data = append(data, x)
	}

	// Insertion Sort
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = key
	}

	// Output array yang sudah diurutkan
	for _, v := range data {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// Cek jarak
	if len(data) < 2 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	jarak := data[1] - data[0]
	tetap := true

	for i := 1; i < len(data)-1; i++ {
		if data[i+1]-data[i] != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}