package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	resx := make([][]uint8, dx)
	for i, _ := range resx {
		rowi := make([]uint8, dy)
		resx[i] = rowi
		for j, _ := range resx[i] {
			resx[i][j] = uint8((i ^ j) * 69)
		}
	}
	return resx
}

// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
func main() {
	pic.Show(Pic)
}
