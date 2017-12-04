package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dx)
	for x := range res {
		res[x] = make([]uint8, dy)
		for y := range res[x] {
			// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
			res[x][y] = uint8((x ^ y) * 69)
		}
	}
	return res
}

func main() {
	pic.Show(Pic)
}
