package main

import (
	"fmt"
	"unicode"
)

func main() {
	var num1, num2 float64
	var operator rune
	var choice rune

	for {
		fmt.Print("Enter first number: ")
		if _, err := fmt.Scanln(&num1); err != nil {
			fmt.Println("Invalid number:", err)
			continue
		}

		fmt.Print("Enter operator (+, -, *, /): ")
		if _, err := fmt.Scanf("%c\n", &operator); err != nil {
			fmt.Println("Invalid operator:", err)
			continue
		}

		fmt.Print("Enter second number: ")
		if _, err := fmt.Scanln(&num2); err != nil {
			fmt.Println("Invalid number:", err)
			continue
		}

		result, err := calculate(num1, num2, operator)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %.2f\n", result)

		fmt.Print("Do you want to calculate more? (y/n): ")
		if _, err := fmt.Scanf("%c\n", &choice); err != nil {
			fmt.Println("Invalid choice:", err)
			continue
		}

		choice = unicode.ToLower(choice)
		if choice == 'n' {
			break
		}

	}
}

func calculate(a, b float64, op rune) (float64, error) {
	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("invalid operator: %c", op)
	}
}
