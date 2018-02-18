package three

import (
	"testing"
)

func TestGenerateGrid(t *testing.T){
	grid := generateGrid(23)
	expected := 5
	if expected != len(grid){
		t.Errorf("Expected: %d Got: %d",expected, len(grid))
	}
	grid = generateGrid(26)
	expected = 7
	if expected != len(grid){
		t.Errorf("Expected: %d Got: %d",expected, len(grid))
	}
}

func TestCalcGridSize(t *testing.T){
	result := calcGridSize(23,0,1)
	expected := 2
	if expected != result {
		t.Errorf("Expected: %d Got: %d",expected, result)
	}
}