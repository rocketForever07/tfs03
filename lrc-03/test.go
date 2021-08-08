package main

import "fmt"
func main(){

	s :="hello"

	s1 :=""

	for _,char := range s{
		s1+=string(char)
	}

	fmt.Println(s1);
}
