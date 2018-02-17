package one

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetTestData(t *testing.T) {
	data, err := getTestData("data/testDataOne.txt")
	fmt.Println(string(data))

	expected := "11221"

	if err != nil {
		t.Errorf("Unable to read testDataOne%s \n", err)
	} else if res := strings.Compare(string(data), expected); res != 0 {
		t.Errorf("Error while reading testData. Got: %s Expected: %s", string(data), expected)
	}

}

func TestGetSum(t *testing.T) {

	data, _ := getTestData("data/input.txt")
	expected := 1251

	if res := getSum(data,1); res != expected {
		t.Errorf("Error while getting sum. Got: %d Expected: %d\n", res, expected)
	}

}

func TestGetSumSecond(t *testing.T){
	data, _ := getTestData("data/input.txt")
	expected := 1244

	if res := getSum(data,2); res != expected {
		t.Errorf("Error while getting sum. Got: %d Expected: %d\n", res, expected)
	}

}
