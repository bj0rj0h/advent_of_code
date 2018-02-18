package two

import (
	"io/ioutil"
	"strconv"
	"math"
	"fmt"
	"strings"
)

func Run(){
	data,_ := getData("2017/two/input.txt")
	slices := bytesToMatrix(data)
	fmt.Printf("TWO-1: Result is: %d \n",calcSumDiff(slices))
	fmt.Printf("TWO-2: Result is: %d \n",getSumOfAllSlices(slices))
}

func getData(path string) ([]byte,error){
	dat, err := ioutil.ReadFile(path)
	checkErr(err)
	return dat, err
}

func bytesToMatrix(bytes []byte) [][]int {
	data := make([][]int, 0)
	inner := make([]int,0)
	var builder strings.Builder
	prevSpace := false
	for _,b := range bytes{
		if b == 32 || b == 9{
			if !prevSpace {
				prevSpace = true
				intS := getIntFromBuilder(&builder)
				inner = append(inner,intS)
				builder.Reset()
			}
		}else if b == 10 {
			prevSpace = false
			intS := getIntFromBuilder(&builder)
			inner = append(inner, intS)
			builder.Reset()
			data = append(data,inner)
			inner = make([]int,0)
		}else {
			prevSpace = false
			builder.WriteByte(b)
		}
	}
	return data
}

func getIntFromBuilder(builder *strings.Builder) int{
	str := builder.String()
	intS,err := strconv.Atoi(str)
	checkErr(err)

	return intS
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

func getSumOfEvenDivsForSlices(slice []int) int{

	var sum int
	for i,val := range slice{
		for j,inner := range slice{
			if i != j{
				if val%inner==0{
					sum += val/inner
				}
			}
		}
	}
	return sum
}

func getSumOfAllSlices(slices [][]int) int{
	var sum int
	for _,slice := range slices{
		sum += getSumOfEvenDivsForSlices(slice)
	}
	return sum

}


func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}