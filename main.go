package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	PrintMessage("ola")
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
2. Função (visão geral)
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

func subtracao(num1 int, num2 int) int { return num1 - num2 }

/*
2.1. retorno de funções
	Os exemplos demonstrados acima mostram a simplicidade e clareza do Go. E acredito que explicito dessa forma,
não precisa de explicação.

3. Os nomes em Go
	Já ouviu falar nisso em outras linguagens e aqui não seria diferente, código limpo, legibilidade e simplicidade
são os principais objetivos do desenvolvimento em Go.
- nomes curtos, concisos e evocativos
- não deve ter a indicação do pacote
- métodos ou funções sem a necessidade de get... Com o tempo vai aprendendo.

TAREFA
Implemente uma função simples que recebe uma mensagem como parâmetro e imprime no console usando
println.
- Crie uma função chamada PrintMessage que recebe um parâmetro message do tipo string
- dentro do corpo da função, coloque o println
*/

func PrintMessage(message string) {
	fmt.Println(message)
}

/*
4. Tipos de dados
	No meu caso, tenho experiências com desenvolvimento e muito mais tempo de estudo, e é estranho
e entediante talvez (não estou achando entediante, não sei o porque) ter que fazer isso. Mas se você está apredendo uma nova linguagem, é necessário. É como aprender
uma nova língua, você precisa saber o alfabeto.
	Go é esteticamente tipado, o compilador precisa saber todos os tipos de dados. E isso é maravilhoso! Os tipos
são:
- primitivos (int, float, bool,string)
- compostos (struct, array, slice, map, channel)
- interfaces
Para declarar uma variável, utilizamos:
*/

var example string

//Ou seja, palavra reservada var, nome da variavel e seu tipo. Claro que podemos tambem colocar de imediato um valor nela:

var example2 string = "exemplo"

/*
4.1. Conversões
	Alguns tipos podem ser convertidos em outros, como por exemplo um valor float para um int.
exemplo seria algo assim:

var number float = 8.8

fmt.Println(int(number))

Lembra a sintaxe do python nesse caso.

5. Constantes e Variáveis
	Esse parte é bem dedutível, a começar com os valores constantes (que não se alteram).

const PI = 3.14

	Aqui eu acabei de aprender que pode sim ter tipagem implícita, que é o caso do exemplo acima.
E tem outra coisa bem interessante é que, se uma variável explícita estiver como parametro de uma função implícita,
essa função se torna tipara.

var value int

func (value + 10)

6. Transbordamento
	As constantes numéricas não têm tamanho e nem sinal. Exemplo:
const Ln2 =  0,693147180559945309417232121458176568075500134360255254120680009

7. Tarefas múltipas e enumeração
	As atribuições podem ser múltipas, exemplos dela:
const carne, veg = "porco" , "soja"
	Com a tipagem fica:
const gato, cachorro, pato string = "Bagdá", "caramelo", "donald"

Bagdá é o nome do meu gato

	Podemos enumerar nossas constantes, exemplo:

const (
	EU = 0
	TU = 1
	ELIX = 2
)

	Também podemos dar um nome e um tipo a nossa enumeração:
type Mouse int (
FIO = 0
SEMFIO = 1
)

PARA SOMENTE AS VARIAVEIS, nós podemos declarar com o operador de atribuição :=, ele é apenas para valores
não tipados explicitamentes.

var num = 5
ou
num := 5

8. Printing

	Quando utilizamos Println com o package fmt, podemos mostrar no console o que bem entedermos, se seguindo as regras.
E quando queremos mostrar um valor atribuído, utilizamos Printf com os %d, %s ou %v. Se vocÊ já teve contato do javascript, estamos
falando disso:

const luz = "luz"
const text = `uma mensagem de ${luz}`

Aqui no GO, faremos da seguinte forma

var luz string = "luz"

fmt.Printf("uma mensagem de %s", luz)


9. Tipos elementares
	Existem 3 tipos elementares em Go:
	- Boolean
	- NUmber
	- Person

9.1. Boolean
	Aí não tem mistério, é true os false (minusculo)

9.2.
*/
