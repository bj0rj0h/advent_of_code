package three

import (
	"fmt"
	"math"
)

func getCoordinatesOf(desiredValue int, grid [][]int) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == desiredValue {
				fmt.Printf("gridval:%d \n", grid[i][j])
				return j, i
			}
		}
	}
	return -1, -1
}

func calculateWalkingDistance(desiredValue int, startVal int, grid [][]int) int {

	if desiredValue == 1 {
		return 0
	}

	qx, qy := getCoordinatesOf(desiredValue, grid)
	fmt.Printf("qx: %d qy: %d\n", qx, qy)
	px, py := (len(grid)-1)/2, (len(grid)-1)/2
	fmt.Printf("px: %d py: %d\n", px, py)

	xSum := math.Abs(float64(px) - float64(qx))
	ySum := math.Abs(float64(py) - float64(qy))

	distance := int(xSum + ySum)

	return distance
}

func fillGrid(grid [][]int, max int, p policy) [][]int {
	fmt.Printf("max:%d\n", max)

	if max == 1 {
		return grid
	}

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

func printGrid(grid [][]int) {
	fmt.Println(len(grid))
	for i := len(grid) - 1; i >= 0; i-- {
		fmt.Printf("%d\n", grid[i])
	}
}
