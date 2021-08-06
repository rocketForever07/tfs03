package main

import (
	
	"fmt"
)


func cal(opr string, a float64, b float64) float64 {

	var result float64
	switch opr {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result=a*b
	case "div":
		if b==0{
			result=-1
		}else{		
			result=a/b
		}
	}

	return result
}

func main() {
	// fmt.Println(cal("div",5,0));

	slice := []int{1,2}
	slice2 := slice

	slice2[0]=2

	fmt.Println(slice2)
	fmt.Println(slice)

}