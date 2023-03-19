package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j <= i; j++ {
			if arr[i] < arr[j] {
				var temp int = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
			fmt.Println(arr)
		}
	}
	return arr
}

func main() {
	arr := []int{5, 2, 3, 1, 6}
	var finalArr []int = bubbleSort(arr)
	fmt.Println(finalArr)
}
