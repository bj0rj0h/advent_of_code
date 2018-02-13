package one

import (
	"fmt"
	"io/ioutil"
)

func Run() {
	fmt.Println("hello from package")
	data, err := getTestData("/one/data/input.txt")
	checkErr(err)
	getSum(data)

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

func nextIsMatch(data []byte) {

}
