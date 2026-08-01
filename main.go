package main

/* # importacoes em varias linhas, fora da convencao, mas valido
import "fmt"
import "github.com/iuribrito/go-learn/lib"
*/

// Multi Import
import (
	"fmt"
	. "fmt" // incorpora as funcoes do pacote no aquivo atual, o uso de dot imports eh desencorajado pela comunidade

	meuFmt "github.com/iuribrito/go-learn/fmt" // alias para pacotes quando nomes complexos ou conflito de nomes
	"github.com/iuribrito/go-learn/lib"        // o pacote eh acessado pelo ultimo nome do importe
	_ "github.com/iuribrito/go-learn/lib"      // inicializa o pacote, mas nao traz para o escopo atual, normalmente utilizado com drivers para banco de dados
	// "github.com/iuribrito/go-learn/lib/internal/foo" # Erro: por se tratar de um pocate interno do lib ele nao pode ser importado fora dele
)

func main() {
	fmt.Println("Hello, Go!")
	Println("Hello, Go!") // funcao incorporada do "fmt"

	fmt.Println(lib.PublicVar)
	// fmt.Println(lib.privateVar) # Erro: variavel privada nao pode ser importada em outro pacote
	meuFmt.Echo("Hello, Go!") //usando a funcao pelo pacote importado com alias

	fmt.Println(somar(1, 2))
	fmt.Println(somarTudo(10, 20, 30, 10))

	y := diminuir(10)(5)
	fmt.Println(y)

	f := diminuir(10)
	x := f(4)
	fmt.Println(x)

	a, b := swap(5, 8)
	fmt.Println(a, b)

	anonima := func() {
		fmt.Println("Hello, anonima!")
	}
	anonima()

	fmt.Println(global)
	variables()
	arrays()
}
