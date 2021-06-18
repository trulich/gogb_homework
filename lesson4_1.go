package main

import (
	"fmt"
)

func sortByInserts(slice []int) []int {
	var j, key int
	for i := range slice {
		j = i -1
		if j >= 0 {
			key = slice[i]
			for slice[j] > key {
				slice[j + 1] = slice[j]
				j -= 1
				if j < 0 {
					break
				}
			}
			slice[j + 1] = key
		}

	}

	return slice

}

func main()  {
	nusm := []int{4,234,4,5,8,42,1}
	fmt.Println(sortByInserts(nusm))
}
