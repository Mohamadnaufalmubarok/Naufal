package main

import "fmt"

func isInside(px, py, cx, cy, r int) bool {
	return (px-cx)*(px-cx)+(py-cy)*(py-cy) <= r*r
}

func main() {
	var x1, y1, r1 int
	var x2, y2, r2 int
	var x, y int

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&x, &y)

	inCircle1 := isInside(x, y, x1, y1, r1)
	inCircle2 := isInside(x, y, x2, y2, r2)

	if inCircle1 && inCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}