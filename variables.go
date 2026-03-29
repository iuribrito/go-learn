package main

import (
	"fmt"
)

var global string = "Global" // Variavel global acessivel em todo o pacote

func variables() {
	var local string = "Local" // Variavel acessecivel no escopo atual

	fmt.Println(local)

	var vInt int
	var vFloat float32
	var vString string
	var vBolean bool

	fmt.Println(vInt, vFloat, vString, vBolean)

	// var nome string
	// var sobrenome string
	// var nome, sobrenome string
	// var nome, sobrenome string = "Iuri", "Brito"
	nome := "Iuri"
	sobrenome := "Brito"

	fmt.Println(nome, sobrenome)

	const num = 10 // untyped constant, assume o tipo do contexto

	takeInt8(num)
	takeInt32(num)
	takeFloat32(num)

	a := 63
	f := float32(a)
	fmt.Println(f)
	s := string(a) // convert string of the rune
	fmt.Println(s)
}

func takeInt8(a int8) {
	fmt.Println(a)
}

func takeInt32(a int32) {
	fmt.Println(a)
}

func takeFloat32(a float32) {
	fmt.Println(a)
}
