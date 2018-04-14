package three

import (
	"fmt"
)

//Run is the main entry point for day 3
func Run() {
	desiredValue := 265149
	grid := generateGrid(desiredValue)
	tmpGrid := fillGrid(grid, 0)
	printGrid(tmpGrid)
	gridMax := len(grid) * len(grid)
	distance := calculateWalkingDistance(desiredValue, gridMax, grid)
	fmt.Printf("THREE-1: Result is:%d\n", distance)
	fmt.Printf("THREE-2: Result is:%d\n", calculateSumOfAdjacents(265149))
}

func generateGrid(max int) [][]int {
	numSquares := calcGridSize(max, 0, 1)
	grid := make([][]int, numSquares*2+1)
	for i := range grid {
		grid[i] = make([]int, numSquares*2+1)
	}
	return grid
}

func calcGridSize(neededMaxVal, gridSize, gridMax int) int {

	if gridMax < neededMaxVal {
		gridMax = 8*(gridSize+1) + gridMax
		return calcGridSize(neededMaxVal, gridSize+1, gridMax)
	}
	return gridSize
}
