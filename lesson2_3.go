package main

import (
	"fmt"
)

// 3. С клавиатуры вводится трехзначное число. 
// Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.

func main()  {
	var someNum int
	fmt.Println("Введите трехзначное число:")
	fmt.Scan(&someNum)

	units :=  someNum%10
	tens := someNum/10%10
	hundreds := someNum/100
	fmt.Println("Единиц: ", units)
	fmt.Println("Десятков", tens)
	fmt.Println("Сотен", hundreds)
	
}
