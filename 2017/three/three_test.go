package three

import (
	"fmt"
	"testing"
)

func TestGenerateGrid(t *testing.T) {
	p := incrementalPolicy{0}
	grid := generateGrid(23, p)
	expected := 5
	if expected != len(grid) {
		t.Errorf("Expected: %d Got: %d", expected, len(grid))
	}
	p = incrementalPolicy{0}
	grid = generateGrid(26, p)
	expected = 7
	if expected != len(grid) {
		t.Errorf("Expected: %d Got: %d", expected, len(grid))
	}
}
func TestCalcGridSize(t *testing.T) {
	result := calcGridSize(23, 0, 1)
	expected := 2
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}
func TestCalculateWalkingDistance_1_PASS(t *testing.T) {
	startVal := 1
	p := incrementalPolicy{0}
	grid := generateGrid(startVal, p)
	printGrid(grid)
	result := calculateWalkingDistance(1, startVal, grid)
	expected := 0
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}

func TestCalculateWalkingDistance_12_PASS(t *testing.T) {
	startVal := 12
	p := incrementalPolicy{0}
	grid := generateGrid(startVal, p)
	result := calculateWalkingDistance(startVal, 25, grid)
	expected := 3
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}
func TestCalculateWalkingDistance_23_PASS(t *testing.T) {
	startVal := 23
	p := incrementalPolicy{0}
	grid := generateGrid(startVal, p)
	gridMax := len(grid) * len(grid)
	result := calculateWalkingDistance(startVal, gridMax, grid)
	expected := 2
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}
func TestCalculateWalkingDistance_16_PASS(t *testing.T) {
	startVal := 16
	p := incrementalPolicy{0}
	grid := generateGrid(startVal, p)
	gridMax := len(grid) * len(grid)
	result := calculateWalkingDistance(startVal, gridMax, grid)
	expected := 3
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}
func TestCalculateWalkingDistance_18_PASS(t *testing.T) {
	startVal := 18
	p := incrementalPolicy{0}
	grid := generateGrid(startVal, p)
	gridMax := len(grid) * len(grid)
	result := calculateWalkingDistance(startVal, gridMax, grid)
	expected := 3
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}
func TestPopulateGrid(t *testing.T) {
	p := incrementalPolicy{0}
	grid := generateGrid(12, p)
	expectedQ1, expectedQ2 := 4, 3
	q1, q2 := firstHalf(12, 25, grid)

	if q1 != expectedQ1 {
		t.Errorf("Expected q1 to be: %d Got: %d", expectedQ1, q1)
	}
	if q2 != expectedQ2 {
		t.Errorf("Expected q2 to be: %d Got: %d", expectedQ2, q2)
	}

	fmt.Println("done")
}
func TestCalculateSumOfAdjacents23_PASS(t *testing.T) {
	maxVal := 23
	result := calculateSumOfAdjacents(maxVal)

	if result != 25 {
		t.Errorf("Expected result to be: %d Got: %d", 25, result)
	}
}
func TestCalculateSumOfAdjacents54_PASS(t *testing.T) {
	maxVal := 54
	result := calculateSumOfAdjacents(maxVal)
	expected := 57
	if result != expected {
		t.Errorf("Expected result to be: %d Got: %d", expected, result)
	}
}
func TestCalculateSumOfAdjacents330_PASS(t *testing.T) {
	maxVal := 330
	result := calculateSumOfAdjacents(maxVal)
	expected := 351
	if result != expected {
		t.Errorf("Expected result to be: %d Got: %d", expected, result)
	}
}
func TestCalculateSumOfAdjacents747_PASS(t *testing.T) {
	maxVal := 747
	result := calculateSumOfAdjacents(maxVal)
	expected := 806
	if result != expected {
		t.Errorf("Expected result to be: %d Got: %d", expected, result)
	}
}
func TestCalculateSumOfAdjacents266330_PASS(t *testing.T) {
	maxVal := 265149
	result := calculateSumOfAdjacents(maxVal)
	expected := 266330
	if result != expected {
		t.Errorf("Expected result to be: %d Got: %d", expected, result)
	}
}

func TestTmp(t *testing.T) {
	desiredValue := 23

	p := incrementalPolicy{0}
	grid := generateGrid(desiredValue, p)
	printGrid(grid)
	fmt.Printf("grid size%d\n", len(grid))
	// fmt.Printf("test:  %d\n", p.getNextValue())
	// p.value++
	// fmt.Printf("test:  %d\n", p.getNextValue())
	tmpGrid := fillGrid(grid, desiredValue, p)
	printGrid(tmpGrid)
	t.Errorf("fail")
}
