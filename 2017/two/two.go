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
	for _,b := range bytes{
		if b == 32 {
			inner = append(inner,s)
			s = ""
		}else if b == 10{
			inner = append(inner,s)
			s = ""
			data = append(data,inner)
			inner = make([]string,0)
		}else {
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