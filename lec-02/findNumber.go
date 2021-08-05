package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func findNumbers(arr []float64) (float64, float64, float64) {

	var sum, avg float64
	count := float64(len(arr))
	max, min := arr[0], arr[0]
	for _, v := range arr {
		sum += float64(v)
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}

	}

	avg = sum / count

	return max, min, avg

}

func main() {

	data, _ := ioutil.ReadFile("./file/numbers.txt")
	s := strings.Split(string(data), " ")
	arr := make([]float64, len(s))

	for i := range s {
		arr[i], _ = strconv.ParseFloat(s[i],64)
	}

	fmt.Println(findNumbers(arr))

}
