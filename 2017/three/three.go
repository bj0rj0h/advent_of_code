package three

import (
	"fmt"
)

type State struct {
	direction rune
	steps     int
}

func Run() {

}

func generateGrid(max int) [][]int {
	numSquares := calcGridSize(max, 0, 1)
	fmt.Printf("NumSquares: %d\n", numSquares)
	grid := make([][]int, numSquares*2+1)
	for i := range grid {
		grid[i] = make([]int, numSquares*2+1)
	}
	return grid
}

func populateGrid(desiredValue int, maxVal int, grid [][]int) (int, int) {
	return firstHalf(desiredValue, maxVal, grid)
}

func firstHalf(desiredValue int, value int, grid [][]int) (int, int) {

	var qx, qy int

	if value == 1 {
		return (len(grid) - 1) / 2, (len(grid) - 1) / 2
	}

	for gridSize := len(grid) - 1; gridSize >= (len(grid)-1)/2; gridSize-- {

		//bottom row
		firstRowIndex := (len(grid) - 1) - gridSize
		bottomRow := grid[firstRowIndex]
		for x := (len(bottomRow) - 1) - firstRowIndex; x > (0 + firstRowIndex); x-- {
			if value == desiredValue {
				qx = x
				qy = firstRowIndex
			}
			bottomRow[x] = value
			value--
		}
		grid[firstRowIndex] = bottomRow
		//leftRow
		for y := 0 + firstRowIndex; y < (len(grid)-1)-firstRowIndex; y++ {
			grid[y][0+firstRowIndex] = value
			if value == desiredValue {
				qx = firstRowIndex
				qy = y
			}
			value--
		}

		//topRow
		topRow := grid[(len(grid)-1)-firstRowIndex]
		for x := 0 + firstRowIndex; x < (len(grid)-1)-firstRowIndex; x++ {
			topRow[x] = value
			if value == desiredValue {
				qx = x
				qy = (len(grid) - 1) - firstRowIndex
			}
			value--
		}

		//rightRow
		for y := (len(grid) - 1) - firstRowIndex; y >= (0 + (len(grid) - 1) - gridSize + 1); y-- {
			grid[y][(len(grid)-1)-firstRowIndex] = value
			if value == desiredValue {
				qx = (len(grid) - 1) - firstRowIndex
				qy = y
			}
			value--
		}

	}

	return qx, qy
}

func secondHalf(grid [][]int) [][]int {
	return grid
}

func calcGridSize(neededMaxVal, gridSize, gridMax int) int {

	if gridMax < neededMaxVal {
		gridMax = 8*(gridSize+1) + gridMax
		return calcGridSize(neededMaxVal, gridSize+1, gridMax)
	}
	return gridSize
}

func calculateWalkingDistance(startVal int, grid [][]int) int {
	//var q1, q2 = getQPos(startVal, grid)

	return 0

}

func getQPos(startVal int, grid [][]int) (int, int) {

	var q1, q2 int

	for i, e := range grid {
		for j, je := range e {
			if je == startVal {
				q1, q2 = i, j
			}
		}
	}

	return q1, q2

}
