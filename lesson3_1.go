package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	var simpleOp bool

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите операцию (+, -, *, /, √, ^, sin, cos): ")
	fmt.Scanln(&op)

	switch op {
	case "√":
		res = math.Sqrt(a)
		simpleOp = true
	case "^":
		res = a * a
		simpleOp = true
	case "sin":
		res = math.Sin(a)
		simpleOp = true
	case "cos":
		res = math.Cos(a)
		simpleOp = true
	}

	if simpleOp == false {

		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)

		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		case "√":
			res = a / b
		default:
			fmt.Println("Операция выбрана неверно")
			os.Exit(1)
		}
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
