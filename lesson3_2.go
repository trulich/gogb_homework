package main

import (
	"fmt"
)

/*
Решил использовать решето Эратосфена, да, самое простой в реализации и самый медленный, но времени не хватает :(
 */

func findNatural(slice []int) (int, []int) {
	var resultSlice []int

	for i := range slice {
		if slice[i]%slice[0] != 0 {
			resultSlice = append(resultSlice, slice[i])
		}
	}
	return slice[0], resultSlice

}

func main() {
	var nums, tmpNum  int
	var s, naturalNums []int

	/*
	Первая итерация, чтобы сгенерить первый массив
	Исходил из того, что известно что 2 это натуральное число
	 */
	naturalNums = append(naturalNums, 2)
	fmt.Print("Введите целое положительное число: ")
	fmt.Scanln(&nums)
	for i := 2; i <= nums; i++ {
		if i%2 != 0 {
			s = append(s, i)
		}
	}
	/*
	Вторая итерация, перебираем массивы (слайсы) и находим натуральные числа
	 */
	for len(s) > 0 {
		tmpNum, s = findNatural(s)
		naturalNums = append(naturalNums, tmpNum)
	}
	fmt.Println("Натуральные числа:\n", naturalNums)
}
