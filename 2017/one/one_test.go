package one

import (
	"testing"
	"fmt"
	"strings"
	)

func TestGetTestData(t *testing.T){
	data,err := getTestData("data/testData.txt")
	fmt.Println(string(data))

	expected := "123425"


	if err != nil{
		t.Errorf("Unable to read testData%s \n",err)
	} else if res := strings.Compare(string(data),expected); res != 0{
		t.Errorf("Error while reading testData. Got: %s Expected: %s",string(data),expected)
	}

}
