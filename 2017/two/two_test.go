package two

import (
	"testing"
	"fmt"
)

func TestGetData(t *testing.T){
	data,_ := getData("testData.txt")
	fmt.Println(data)
}

func TestBytesToSlice(t *testing.T){
	data,_ := getData("input.txt")
	fmt.Println(data)
	slice := bytesToSlice(data)
	fmt.Println(slice)
}



