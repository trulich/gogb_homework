package main

import (
	"fmt"
	"math"
)

func getCyrcleD(s float64) float64  {
	return 2 * math.Sqrt(s/math.Pi)
}

func getCyrcleC(d float64) float64  {
	return math.Pi * d
}

// 2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. 
// Площадь круга должна вводиться пользователем с клавиатуры.

func main()  {
	var cyrcleS, cyrcleD, cyrcleC float64
	
	fmt.Println("Введите площадь окружности: ")
	fmt.Scan(&cyrcleS)
	
	cyrcleD = getCyrcleD(cyrcleS)
	cyrcleC = getCyrcleC(cyrcleD)

	fmt.Println("Диаметр окружности:", cyrcleD, "\nДлина окружности:", cyrcleC)
}
