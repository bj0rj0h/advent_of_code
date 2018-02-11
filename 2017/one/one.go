package one

import (
	"fmt"
	"io/ioutil"
)

func Run(){
	fmt.Println("hello from package")
	//getTestData("")

}

func getTestData(path string) ([]byte,error){
	dat, err := ioutil.ReadFile(path)
	checkErr(err)

	return dat,err
}

func checkErr(e error){
	if e != nil{
		fmt.Println(e)
	}
}
