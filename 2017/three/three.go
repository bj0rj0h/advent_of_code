package three

import (
	"fmt"
)

type State struct {
	direction rune
	steps int
}

func Run(){

}

func generateGrid(max int)[][]int{
	numSquares := calcGridSize(max,0,1)
	fmt.Println("NumSquares: %d",numSquares)
	return make([][]int, numSquares*2+1)
}

func calcGridSize(neededMaxVal,gridSize,gridMax int) int{

	if gridMax < neededMaxVal {
		gridMax = 8*(gridSize+1)+gridMax
		return calcGridSize(neededMaxVal,gridSize+1,gridMax)
	}
		return gridSize
}

func calculateWalkingDistance(startVal int, grid [][]int) int{



	return 0
}