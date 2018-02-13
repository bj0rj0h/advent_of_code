package one

import (
	"fmt"
	"io/ioutil"
)

func Run() {
	data, err := getTestData("2017/one/data/input.txt")
	checkErr(err)
	result := getSum(data)
	fmt.Printf("ONE-1: Result is %d: \n", result)
	result = getSumSecond(data)
	fmt.Printf("ONE-2: Result is %d: \n", result)

}
func getSum(bytes []byte) int {

	stringData := string(bytes)
	var sum int
	var cmp int
	for i, r := range stringData {
		if i != len(bytes)-1 {
			cmp = int(stringData[i+1] - '0')
		} else {
			cmp = int(stringData[0] - '0')
		}
		if int(r-'0') == cmp {
			sum += int(r - '0')
		}

	}

	return sum
}

func getSumSecond(bytes []byte) int {

	stringData := string(bytes)
	var sum int
	var cmp int
	for i, r := range stringData {
		if i > len(stringData)/2-1 {
			x := i-len(stringData)/2
			cmp = int(stringData[0+x] - '0')
		} else {
			cmp = int(stringData[i+len(stringData)/2] - '0')
		}

		if int(r-'0') == cmp {
			sum += int(r - '0')
		}
	}

	return sum
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
