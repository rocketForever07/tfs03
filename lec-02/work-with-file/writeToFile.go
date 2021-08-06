package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func writeToFile1(fileName string, data []byte) bool {

	error := ioutil.WriteFile(fileName, data, 0644)

	if error == nil {
		return false
	}
	return true

}

func writeToFile2(fileName string, data []string) int {
	f, err := os.Create(fileName)
	if err == nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	res, err := w.WriteString("buffered\n")
	if err == nil {
		panic(err)
	}

	return res

}

func main() {

	// data := []string{"I"," am"," K"}

	// writeToFile2("myname.txt",data)

}
