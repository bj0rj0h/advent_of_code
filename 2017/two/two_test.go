package two

import (
	"testing"
)

func TestGetData(t *testing.T){
	data,_ := getData("testData.txt")
	var sum int
	expected := 1126
	for _,val := range data{
		sum += int(val)
	}
	if sum != expected{
		t.Errorf("Wrong data read from file. Got: %d expected: %d\n",sum,expected)
	}
}

func TestBytesToSlice(t *testing.T){
	data,_ := getData("testData.txt")
	slices := bytesToMatrix(data)
	var sum int
	expected := 198
	for _,slice := range slices{
		for _,s := range slice{
			sum += s
		}
	}
	if sum != expected{
		t.Errorf("Wrong data read from file. Got: %d expected: %d\n",sum,expected)
	}
}

func TestCalcDiff(t *testing.T){
	data,_ := getData("testData.txt")
	slices := bytesToMatrix(data)
	expected := 6
	if diff := calcDiff(slices[1]); diff != expected{
		t.Errorf("Error while getting diff. Got %d expected %d\n",diff,expected)
	}
}

func TestCalcSumOfDiffs(t *testing.T){
	data,_ := getData("testData.txt")
	slices := bytesToMatrix(data)
	expected := 101
	if diff := calcSumDiff(slices); diff != expected{
		t.Errorf("Error while getting diff. Got %d expected %d\n",diff,expected)
	}
}

func TestGetSumOfEvenDivs(t *testing.T){
	slice := []int{5,9,2,8}
	result := getSumOfEvenDivsForSlices(slice)
	expected := 4
	if result != expected{
		t.Errorf("Got: %d Expected: %d\n",result,expected)
	}
}

func TestGetSumOfAllSlices(t *testing.T){
	data,_ := getData("input.txt")
	slices := bytesToMatrix(data)
	result := getSumOfAllSlices(slices)
	expected := 214
	if result != expected{
		t.Errorf("Got: %d Expected: %d\n",result,expected)
	}
}