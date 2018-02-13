package one

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetTestData(t *testing.T) {
	data, err := getTestData("data/testDataOne.txt")
	fmt.Println(string(data))

	expected := "1122"

	if err != nil {
		t.Errorf("Unable to read testDataOne%s \n", err)
	} else if res := strings.Compare(string(data), expected); res != 0 {
		t.Errorf("Error while reading testData. Got: %s Expected: %s", string(data), expected)
	}

}

func TestGetSum(t *testing.T) {

	data, _ := getTestData("data/testDataOne.txt")
	expected := 4

	if res := getSum(data); res != expected {
		t.Errorf("Error while getting sum. Got: %d Expected: %d\n", res, expected)
	}

}
