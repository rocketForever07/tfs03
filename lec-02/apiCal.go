package main

import (
	
	"fmt"
	"net/http"
)


type DataRes struct {

	a float64
	b float64
	result float64
}

func cal(w http.ResponseWriter, a float) float64 {

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

http.Handle("/",)


}