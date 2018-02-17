package two

import (
	"testing"
)

func TestGetData(t *testing.T){
	data,_ := getData("testData.txt")
	var sum int
	expected := 1158
	for _,val := range data{
		sum += int(val)
	}
	if sum != expected{
		t.Errorf("Wrong data read from file. Got: %d expected: %d",sum,expected)
	}
}

func TestBytesToSlice(t *testing.T){
	data,_ := getData("testData.txt")
	slices := bytesToSlice(data)
	var sum int
	expected := 198
	for _,slice := range slices{
		for _,s := range slice{
			sum += s
		}
	}
	if sum != expected{
		t.Errorf("Wrong data read from file. Got: %d expected: %d",sum,expected)
	}
}

func TestCalcDiff(t *testing.T){
	/*diff := calcDiff(slice []string){

	}*/
}





