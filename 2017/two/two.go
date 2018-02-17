package two

import (
	"io/ioutil"
	"strconv"
)

func Run(){

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

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}