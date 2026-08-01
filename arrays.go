package main

import (
	"fmt"
)

func arrays() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1: 2, 2: 3}

	fmt.Println(arr1)
	fmt.Println(arr2)

	i := 10
	// for i := 0; i < 10; i++ {
	// for ;i < 10; i++ {
	// for { # loop infinito
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := range 10 {
		fmt.Println(i)
	}

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, el := range arr {
		fmt.Println(i, el)
	}
}
