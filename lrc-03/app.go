package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Exps struct {
	expression string
}

/*
//implement simple stack from array(push,pop,peek)
*/
type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Remove and return top element of stack.
func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

// return top element of stack.

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element
	}
}

/*
//end
*/

//implement theo huong dan https://www.tutorialspoint.com/Convert-Infix-to-Postfix-Expression
func getPriority(op string) int {
	if op == "x" || op == "/" || op == "%" {
		return 2
	} else if op == "+" || op == "-" {
		return 1
	} else {
		return 0
	}
}
func isOpr(s rune) bool {
	if string(s) == "/" || string(s) == "x" || string(s) == "-" || string(s) == "+" {
		return true
	} else {
		return false
	}
}
func infixToPostfix(infix string) string {

	var stack Stack
	postfix := ""

	for _, char := range infix {
		if isOpr(char) == false {
			postfix += string(char)
		} else if string(char) == "(" {
			stack.Push(string(char))
		} else if string(char) == ")" {
			var x string = stack.Pop()
			for string(x) != "(" {
				postfix += string(x)
				x = stack.Pop()
			}
		} else {
			for stack.IsEmpty() == false && getPriority(string(char)) <= getPriority(stack.Peek()) {
				postfix += string(stack.Pop())
			}
			stack.Push(string(char))
		}
	}

	for stack.IsEmpty() == false {
		postfix += stack.Pop()
	}

	return postfix
}
func excute(expr string) string {
	postfix := infixToPostfix(expr)
	var stack Stack

	for _, char := range postfix {

		if isOpr(char) == false {
			stack.Push(string(char))
		} else {
			x, _ := strconv.ParseFloat(stack.Pop(), 64)
			y, _ := strconv.ParseFloat(stack.Pop(), 64)

			switch string(char) {
			case "+":
				y += x
			case "-":
				y -= x
			case "x":
				y *= x
			case "/":
				y /= x
			}
			stack.Push(strconv.FormatFloat(y, 'f', 6, 64))
		}
	}

	return stack.Pop()
}

//end

func cal(w http.ResponseWriter, r *http.Request) {
	var exps Exps

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	decoder.Decode(exps)

	result := excute(exps.expression)
	res := Exps{expression: string(result)}
	encoder.Encode(res)

}

func main() {

	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/cal", cal)
	server.ListenAndServe()

}
