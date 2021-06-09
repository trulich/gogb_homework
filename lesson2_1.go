package main

import (
	"fmt"
)

func rectangleArea(a, b int) int  {
	return a * b
}

// 1. Напишите программу для вычисления площади прямоугольника. 
// Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.


func main()  {
	var rectangleLen, rectangleWidth int
	
	fmt.Println("Введите длину прямоугольника: ")
	fmt.Scan(&rectangleLen)
	fmt.Println("Введите ширину прямоугольника: ")
	fmt.Scan(&rectangleWidth)	
	
	fmt.Println("Площадь прямоугольника:", rectangleArea(rectangleLen, rectangleWidth))
}
