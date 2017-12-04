package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dx)
	for i := range res {
		res[i] = make([]uint8, dy)
		for j := range res[i] {
			// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
			res[i][j] = uint8((i ^ j) * 69)
		}
	}
	return res
}

func main() {
	pic.Show(Pic)
}
