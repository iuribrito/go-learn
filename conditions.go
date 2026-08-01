package main

import (
	"fmt"
	"math"
	"time"
)

func conditions() {
	y := 17
	if y < 18 {
		fmt.Println("menor de idade")
	} else if y >= 18 && y < 60 {
		fmt.Println("maior de idade")
	} else {
		fmt.Println("idoso")
	}

	if x := 10; x > 5 {
		fmt.Println("x maior que 5")
	}

	z := 1
	switch z {
	case 1:
		fmt.Println(1)
		//fallthrough # executa tbm o proximo case
	case 2, 3:
		fmt.Println("2 ou 3")
	default:
		fmt.Println("default")
	}

	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("resultado 2")
	default:
		fmt.Println("resultado anormal")
	}

	t := time.Now()
	switch {
	case t.Weekday() > 0 && t.Weekday() < 6:
		fmt.Println("week")
	default:
		fmt.Println("weekend")
	}
}
