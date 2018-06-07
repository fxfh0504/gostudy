package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8,dy,dy)
	for y:=0; y<dy;y++  {
		picx := make([]uint8,dx,dx)
		for x:=0;x<dx;x++ {
			picx[x] =uint8(x + y)
		}
		pic[y] =picx

	}
	return pic
}

func main() {
	//uint8s := Pic(10, 15)
	pic.Show(Pic)
}
