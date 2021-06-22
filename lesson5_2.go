package main

import "fmt"
// кажется я не понял как оптимизировать с помощью мап
func fibonacchi() func() int {
	num  := map[string]int{
		"first": 1,
		"second": 2,
	}
		return func() int {
		res := num["first"]
		num["first"], num["second"] =  num["second"], num["first"] + num["second"]
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