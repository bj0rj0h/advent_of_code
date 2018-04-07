package three

import (
	"fmt"
	"testing"
)

func TestGenerateGrid(t *testing.T) {
	grid := generateGrid(23)
	expected := 5
	if expected != len(grid) {
		t.Errorf("Expected: %d Got: %d", expected, len(grid))
	}
	grid = generateGrid(26)
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
	grid := generateGrid(startVal)
	result := calculateWalkingDistance(1, startVal, grid)
	expected := 0
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}

func TestCalculateWalkingDistance_12_PASS(t *testing.T) {
	startVal := 12
	grid := generateGrid(startVal)
	result := calculateWalkingDistance(startVal, 25, grid)
	expected := 3
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}

func TestCalculateWalkingDistance_23_PASS(t *testing.T) {
	startVal := 23
	grid := generateGrid(startVal)
	gridMax := len(grid) * len(grid)
	result := calculateWalkingDistance(startVal, gridMax, grid)
	expected := 2
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}
func TestCalculateWalkingDistance_16_PASS(t *testing.T) {
	startVal := 16
	grid := generateGrid(startVal)
	gridMax := len(grid) * len(grid)
	result := calculateWalkingDistance(startVal, gridMax, grid)
	expected := 3
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}
func TestCalculateWalkingDistance_18_PASS(t *testing.T) {
	startVal := 18
	grid := generateGrid(startVal)
	gridMax := len(grid) * len(grid)
	result := calculateWalkingDistance(startVal, gridMax, grid)
	expected := 3
	if expected != result {
		t.Errorf("Expected: %d Got: %d", expected, result)
	}
}
func TestPopulateGrid(t *testing.T) {
	grid := generateGrid(12)
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
