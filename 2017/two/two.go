package two

import (
	"io/ioutil"
	"fmt"
)

func Run(){

}

func getData(path string) ([]byte,error){
	dat, err := ioutil.ReadFile(path)
	checkErr(err)
	return dat, err
}

func bytesToSlice(bytes []byte) [][]string {
	data := make([][]string, 0)
	inner := make([]string,0)
	s := ""
	prevSpace := false
	for _,b := range bytes{
		if b == 32 || b == 9{
			if !prevSpace {
				prevSpace = true
				inner = append(inner, s)
				s = ""
			}
		}else if b == 10 {
			prevSpace = false
			inner = append(inner,s)
			s = ""
			data = append(data,inner)
			inner = make([]string,0)

		}else {
			prevSpace = false
			s += string(b)
		}

	}
	return data
}

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}