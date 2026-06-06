package main

import "fmt"

func main() {
	var num float64
	var sum float64
	var count int

	for {
		fmt.Scan(&num)
		
		if num == 9999 {
			break
		}
		
		sum += num
		count++
	}

	if count > 0 {
		rerata := sum / float64(count)
		fmt.Println(rerata)
	} else {
		fmt.Println(0)
	}
}