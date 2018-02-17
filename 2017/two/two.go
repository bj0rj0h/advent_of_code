package two

import (
	"io/ioutil"
	"strconv"
	"math"
	"fmt"
)

func Run(){
	data,_ := getData("2017/two/input.txt")
	slices := bytesToSlice(data)
	fmt.Println(calcSumDiff(slices))
}

func getData(path string) ([]byte,error){
	dat, err := ioutil.ReadFile(path)
	checkErr(err)
	return dat, err
}

func bytesToSlice(bytes []byte) [][]int {
	data := make([][]int, 0)
	inner := make([]int,0)
	s := ""
	prevSpace := false
	for _,b := range bytes{
		if b == 32 || b == 9{
			if !prevSpace {
				prevSpace = true
				intS,err := strconv.Atoi(s)
				checkErr(err)
				inner = append(inner, intS)
				s = ""
			}
		}else if b == 10 {
			prevSpace = false
			intS,err := strconv.Atoi(s)
			checkErr(err)
			inner = append(inner, intS)
			s = ""
			data = append(data,inner)
			inner = make([]int,0)
		}else {
			prevSpace = false
			s += string(b)
		}
	}
	return data
}

func calcDiff(slice []int) int {
	high := math.MinInt64
	low := math.MaxInt64
	for _,val := range slice{
		if val < low{
			low = val
		}
		if val > high{
			high = val
		}
	}
	return high - low
}

func calcSumDiff(ints [][]int) int {
	var sum int
	for _,slice:= range ints{
		sum += calcDiff(slice)
	}
	return sum
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}