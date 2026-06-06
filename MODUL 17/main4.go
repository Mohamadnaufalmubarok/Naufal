package main

import "fmt"

func main() {
	var num int
	var max, min int
	pertama := true

	for {
		fmt.Scan(&num)

		if num == -1 {
			break
		}

		if pertama {
			max = num
			min = num
			pertama = false
		} else {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
	}

	if !pertama {
		fmt.Printf("a. Maksimum: %d\n", max)
		fmt.Printf("b. Minimum: %d\n", min)
	} else {
		fmt.Println("Tidak ada data yang dimasukkan")
	}
}