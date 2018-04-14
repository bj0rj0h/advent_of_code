package three

import (
	"fmt"
	"math"
)

func calculateWalkingDistance(desiredValue int, startVal int, grid [][]int) int {

	qx, qy := firstHalf(desiredValue, startVal, grid)
	px, py := (len(grid)-1)/2, (len(grid)-1)/2

	xSum := math.Abs(float64(px) - float64(qx))
	ySum := math.Abs(float64(py) - float64(qy))

	distance := int(xSum + ySum)

	return distance
}

func fillGrid(grid [][]int, max int, p policy) [][]int {
	fmt.Printf("max:%d\n", max)

	gridSize := 1
	x := (len(grid)) / 2
	y := (len(grid)) / 2
	p = p.getNextValue()
	grid[y][x] = p.getValue()
	x = x + 1
	p = p.getNextValue()
	grid[y][x] = p.getValue()
	for range grid {
		//UP
		for i := 0; i < (1 + (2 * (gridSize - 1))); i++ {
			y++
			fmt.Printf("UP y: %d x: %d\n", y, x)
			p = p.getNextValue()
			grid[y][x] = p.getValue()
			fmt.Printf("p.getNextValue():%d\n", grid[y][x])
			if grid[y][x] == max {
				return grid
			}
		}

		//LEFT
		for i := 0; i < (2 * gridSize); i++ {
			x--
			fmt.Printf("LEFT y: %d x: %d\n", y, x)
			p = p.getNextValue()
			grid[y][x] = p.getValue()
			fmt.Printf("count:%d\n", grid[y][x])
			if grid[y][x] == max {
				return grid
			}
		}

		//DOWN
		for i := 0; i < (2 * gridSize); i++ {
			y--
			fmt.Printf("DOWN y: %d x: %d\n", y, x)
			p = p.getNextValue()
			grid[y][x] = p.getValue()
			fmt.Printf("count:%d\n", grid[y][x])
			if grid[y][x] == max {
				return grid
			}

		}

		//RIGHT
		for i := 0; i < ((2 * gridSize) + 1); i++ {
			x++
			fmt.Printf("RIGHT y: %d x: %d\n", y, x)
			p = p.getNextValue()
			grid[y][x] = p.getValue()
			fmt.Printf("count:%d\n", grid[y][x])
			if grid[y][x] == max {
				return grid
			}
		}
		gridSize++
	}
	return grid
}

func calculateSumOfAdjacents(maxVal int) int {

	grid := make([][]int, 101)
	gridSize := 1
	for i := range grid {
		grid[i] = make([]int, 101)
	}
	x := 50
	y := 50
	grid[y][x] = 1
	x = x + 1
	grid[y][x] = 1
	for {
		//UP
		for i := 0; i < (1 + (2 * (gridSize - 1))); i++ {
			y++
			grid[y][x] = getSumOfAdjacents(grid, x, y)
			if grid[y][x] > maxVal {
				return grid[y][x]
			}
		}

		//LEFT
		for i := 0; i < (2 * gridSize); i++ {
			x--
			grid[y][x] = getSumOfAdjacents(grid, x, y)
			if grid[y][x] > maxVal {
				return grid[y][x]
			}
		}

		//DOWN
		for i := 0; i < (2 * gridSize); i++ {
			y--
			grid[y][x] = getSumOfAdjacents(grid, x, y)
			if grid[y][x] > maxVal {
				return grid[y][x]
			}
		}

		//RIGHT
		for i := 0; i < ((2 * gridSize) + 1); i++ {
			x++
			grid[y][x] = getSumOfAdjacents(grid, x, y)
			if grid[y][x] > maxVal {
				return grid[y][x]
			}
		}
		gridSize++
	}
}

func getSumOfAdjacents(grid [][]int, x, y int) int {
	sum := 0

	sum += grid[y][x+1]
	sum += grid[y+1][x+1]
	sum += grid[y+1][x]
	sum += grid[y+1][x-1]
	sum += grid[y][x-1]
	sum += grid[y-1][x-1]
	sum += grid[y-1][x]
	sum += grid[y-1][x+1]

	return sum
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

func printGrid(grid [][]int) {
	fmt.Println(len(grid))
	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Printf("%d\n", grid[i])
	}
}
