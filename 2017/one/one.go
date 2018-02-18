package one

import (
	"fmt"
	"io/ioutil"
)

func Run() {
	data, err := getTestData("2017/one/data/input.txt")
	checkErr(err)
	result := getSum(data,1)
	fmt.Printf("ONE-1: Result is: %d \n", result)
	result = getSum(data,2)
	fmt.Printf("ONE-2: Result is: %d \n", result)

}
func getSum(bytes []byte,challenge int) int {

	stringData := string(bytes)
	var sum int
	var cmp int
	for i, r := range stringData {

		if challenge == 1{
			cmp = cmpOne(i,stringData)
		}else if challenge == 2{
			cmp = cmpTwo(i,stringData)
		}
		if int(r-'0') == cmp {
			sum += int(r - '0')
		}
	}
	return sum
}


func cmpOne(i int, stringData string) int{

	if i != len(stringData)-1 {
		return int(stringData[i+1] - '0')
	} else {
		return int(stringData[0] - '0')
	}
}

func cmpTwo(i int, stringData string) int{
	if i > len(stringData)/2-1 {
		x := i-len(stringData)/2
		return int(stringData[0+x] - '0')
	} else {
		return int(stringData[i+len(stringData)/2] - '0')
	}
}

func getTestData(path string) ([]byte, error) {
	dat, err := ioutil.ReadFile(path)
	checkErr(err)

	return dat, err
}

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
