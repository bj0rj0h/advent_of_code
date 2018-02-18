package three

import (
	"fmt"

)

func Run(){

}

func generateGrid(max int)[][]int{
	numSquares := calcGridSize(max)
	fmt.Println("NumSquares: %d",numSquares)
	return make([][]int, numSquares*2+1)
}

func calcGridSize(max int) int{



	gridMax := 1
	squares := 0
	for gridMax < max{
		squares += 1
		gridMax = 8*squares+gridMax
	}

	return squares

}