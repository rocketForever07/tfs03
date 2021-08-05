package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func contain(strRoot []string, items []string) bool {

	strMap := make(map[string]int)

	for _, v := range strRoot {
		strMap[v] = 1
	}

	for _, v := range items {
		_, ok := strMap[v] //ok = true nếu map có key=v
		if ok == false {
			return false
		}
	}
	return true
}

func main() {

	data, _ := ioutil.ReadFile("./file/contain.txt")

	strRoot := strings.Split(string(data), " ")

	items := []string{"khoi", "thuy"}

	fmt.Println(contain(strRoot, items))

}
