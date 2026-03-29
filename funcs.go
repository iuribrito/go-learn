package main

func somar(a int, b int) int {
	sum := a + b
	return sum
}

// argumento variádico, deve ser sempre o ultimo argumento da funcao
func somarTudo(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}

func diminuir(a int) func(int) int {
	return func(b int) int {
		return a - b
	}
}

func swap(a, b int) (int, int) {
	return b, a
}

func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return // naked return, retorna as variaveis definida na assinatura de retorno, nao muito encorajado essa pratica
}
