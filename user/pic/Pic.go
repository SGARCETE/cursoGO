package main

import (
	"golang.org/x/tour/pic"
)


func Pic(dx, dy int) [][]uint8 {
	matriz:=  make([][]uint8,dy)

	for i:=range matriz{
		matriz[i]= make([]uint8, dx)
		for j:=range matriz[i]{
			matriz[i][j] = uint8(j*i)
		}
	}
	return matriz
}



func main() {
	pic.Show(Pic)
}
