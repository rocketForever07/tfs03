package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func isPrime(n float64) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else {
		// var i int64 =2
		for i := 2; i <= int(math.Sqrt(n)); i++ {
			if int(n)%i == 0 {
				return true
			}
		}

		return false
	}
}

func main() {
	data, err := ioutil.ReadFile("./file/numbers.txt")
	if err == nil {
		fmt.Println(err)
	}
	s := strings.Split(string(data), " ")

	fmt.Println(s)
	arr := make([]float64, len(s))

	for i := range s {
		arr[i], _ = strconv.ParseFloat(s[i], 64)
		fmt.Println(isPrime(arr[i]))
	}
	
	fmt.Println(arr)

}
