package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}
}

/*
1. Funcionalidade de importação
	Uma lib, módulo ou namespace é chamado de pacote e aqui não é diferente. Cada arquivo GO
pertence a apenas um pacote, enquanto um pacote pode ter vários arquivos Go. Exemplo, se o
seu arquivo cógido pertence a um pacote chamado main (como acima), basta escrever a primeira
linha. Todo aplicativo Go contém uma pasta princiṕal.
	Para compilar um programa, os pacotes e arquivos devem ser compilados na ordem correta, e as
depedências determinam a ordem de compilação dos pacotes, isso significa que se um pacote for
alterado e recompilado, todos os programas também deverão ser recompilados.

1.1. Import
	Um programa em Go é criado vinculando um conjunto de pacotes com a palavra import. Se
quisermos importar um pacote, basta olhar a linha 3, nessa caso o pacote "fmt" informa ao
Go que o programa precisa de funções.
	Os nomes dos pacotes são colocados entre aspas, e como a importação carrega as declaraçoes
públicas do pacote compilado, temos uma forma elegante de fazer isso, que é olhando a linha 3.

1.2. Regras de visibilidade
	Os pacotes expõem seus objetos de código ao código externo com a regra de, quando um
identificador (constante, variável, tipo, função...) começa com a letra maiuscula, então o
objeto é visível no código fora do pacote e é considerado exportado. Exemplo:
	"pack1 . Objeto"
Neste caso, estamos importando Object (se utiliza a notação de ponto, comum em OOP). Também
pode receber outro nome ao usar no código da importação
	import fm "fmt"
*/

func functionName() {
	println("It's just an example")
}

/*
2. Função
	A mais simples declaração de função segue exemplo na linha 50. Fornecer parâmetros é opcional
mas, se caso tenha é necessário 'tipar'.
	A função main como ponto de partida é necessário, senão teremos erro de compilação e em Go, não
precisamos colocar argumentos ou retorno na função main, ela só precisa estar lá. E o programa encerra
imediatamente com sucesso quando main retorna. Percebe-se que o corpo de função é colocado entre chaves.

*/

func numbers(num1 int, num2 int) {
	println(num1 * num2)
}

func soma(num1 int, num2 int) int {
	return num1 + num2
}

func subtracao(num1 int, num2 int) int {}
