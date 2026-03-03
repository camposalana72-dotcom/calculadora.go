package main

import "fmt"

func main(){

	fmt.Println("Enter which first number:")

	var num1,num2 float64
	var operation string
    fmt.Scan(&num1)

	fmt.Println("Enter which second number: ")
	fmt.Scan(&num2)
	
	fmt.Println("Enter which operation: ")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		fmt.Println("Result : ", num1, "+", num2, "=", num1 + num2)
	case "-":
		fmt.Println("Result : ", num1, "-", num2, "=", num1 - num2)
		
	case "*":
		fmt.Println("Result : ", num1, "*", num2, "=", num1 * num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Result : ", num1, "/", num2, "=", num1 / num2)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	}
}



