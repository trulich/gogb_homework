package main

import "fmt"

func fibonacchi() func() int {
	a, b := 0, 1
	return func() int {
		res := a
		a, b = b, a + b
		return  res
	}
}

func main()  {
	var a int

	fmt.Println("Введите целое положительное число: ")
	fmt.Scanln(&a)
	z := fibonacchi()
	for i :=0; i <= a; i++ {
		fmt.Println(z())
	}


}