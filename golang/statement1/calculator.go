package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operator:%s", op)
	}
}
