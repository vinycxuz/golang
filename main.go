package main

import (
	"awesomeProject/DataStructures"
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello World")
	log.Println(req.URL.Path)
	fmt.Fprint(w, "oi"+req.URL.Path[1:])
}

func main() {
	cat := new(DataStructures.List)
	cat.InserirInicio(3)
	cat.InserirInicio(3)
	fmt.Println(cat.Busca(3))
	cat.Imprimir()

	dog := new(DataStructures.List2)
	dog.InserirInicio2(3)
	dog.InserirInicio2(4)
	dog.InserirFinal2(2)
	dog.ImprimirCrescente()
	dog.ImprimirDecrescente()
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

9.2. Number
	Em GO, existe algo interessante na tipagem numerica, nós podemos colocar o comprimento apropriado
para a máquina em que ele é executado. Exemplo:
	int (32 ou 64, depende da maquina), porém podemos definir por exemplo um int32

var a int
var b int32

	E não é permitido a soma entre eles.

9.2.1. Numeros aleatórios
	Dentro do Go, existe o pacote "math/rand", que implementa geradores de pseudoaleatórios

a := rand.Int()
b := rand.Intn(8) --- Aqui ele gera um aleatorio de 0 ate n

9.3. Pacote unicode
	Existem algumas funções uteis dentro do pacote para testar caracteres. Exemplos:
unicode.IsLetter(ch)
unicode.IsDigit(ch)
unicode.IsSpace(ch)

Essas funções retornam bool e o pacote utf8 tambem contem funções para trabalhar com runes.

9.4. Operadores básicos
	Os operadores são os mesmos na qual já estou habituado, então, irei apenas coloca-los aqui para continuar meu passo a passo
== igual
!= nao igual
<
>
<=
>=
...
9.5. Bitwaise operators
	Bitwise são operadores bit a bit que só funcionam com padrões de bits de tamanhos iguais.
Eles são:
& AND
ou OR
^ XOR

Existe mais um que é extremamente interessante:

<< Operador de deslocamento para a esquerda
>> Operador de deslocamento para a direita

    x := 3
    y := 5
    z := (x ^ y) & (x << 2)
    fmt.Println(z)

Você sabe quanto da esse resultado? Para resolver, basta lembrar qe tudo está em binário e que "é XOR e não elevado ao quadrado.

*** Challenge ***

Implementar a função to Fahrenheit( t Celsius) para conversar C para F. Os tipos são float32.
A formula: F = Celsius * 9/5) + 32


*/

func toFahrenheit(Celsius float32) float32 { return (Celsius * 9 / 5) + 32 }

/*
10. Strings
	String são uma sequencia de caracters UTF-8. Quando é declarado, reserva automaticamente 4bytes, mas o Go é inteligente
o suficiente para reservar apenas one-nyte se a string tiver apenas 1 ASCII character.

Vantagens da string em GO é que ocupam menos espaço na memória e como UTF-8 é padrão, Go não precisa codificar e decodificar
strings como em outras linguagens. Vale lembrar que são matrizes imutáveis de bytes.

Existem dois tipos literais de string:
- Interpreted String
- Raw String

esse se tiver duvida, é bom consultar a documentação.

	E ler o tamanho das strings em Go, é algo semelhante a Javascript (senão me engano). Nós simplesmente colocamos len(str).
Por exemplo:
var str string = "Eu"

fmt.Println(len(str))

Apesar disso, vale ressaltar que não podemos pegar o endereço de uma das partes da string como se ela fosse um array.

10.1. Strings
	Toda linguagem existe suas formas pré-definidas de manipular as strings. Vamos mostrar aqui em Go com o package strings e strconv.
a) Prefixos e Sufixos
	strings.HasPrefix(s, prefix string) bool
	strings.HasSuffix(s, suffix string) bool

	precisamos importar o package strings e dentro do método, o primeiro é qual prefixo queremos e o segundo, qual a string
para ser analisada. Mesma lógica com sufixo.

b) Contains
	função boolean que retorna se contem uma substring na string.
	strings.Contains(str, substr string) bool

c) Replace substring
	Podemos recolocar em uma substring uma nova string. Muito bom. Funciona assim:
strings.Replace(str, old, new string, n int)

d) Counting
	Podemos também contar o número de ocorrências em uma substring.
strings.Count(s, str string) int

var str string = "Hello, how is it going, Hugo?"
    var manyG = "gggggggggg"
    fmt.Printf("Number of H's in %s is: ", str)
    fmt.Printf("%d\n", strings.Count(str, "H"))         // count occurences
    fmt.Printf("Number of double g's in %s is: ", manyG)
    fmt.Printf("%d\n", strings.Count(manyG, "gg"))

e) Repeat
	Podemos repetir uma string como se fosse um laço for:
strings.Repeat(s, count int) string

var origS string = "Hi there! "
    var newS string
    newS = strings.Repeat(origS, 3)     // Repeating origS 3 times
    fmt.Printf("The new repeated string is: %s\n", newS)

11. Ponteiros
	Aqui é maravilhoso, e acredito que tenha sido um dos principais motivos de eu estar migrando para Go. Umas das coisas
mais legais que já estudei na vida foi ponteiros, isso em C. E agora veremos como funciona por aqui.
	Em Go, nós podemos definir se uma variável é um ponteiro ou não. Isso permite ao usuário controle básico sobre o layout básico
da mémoria, o controle do tamanho total de uma determinada coleção de estrutura de dados, numero de alocações e os padrões de acesso
à memória.
	Assim como em C, ao colocarmos o operador &, temos o endereço de memória daquela variável. Exemplo
var oi = 5
fmt. Printf("Isso é um numero %d, e está localizado aqui %p", oi, &oi)

Para declarar uma variável como um ponteiro, faremos da seguinte forma
var example * int
example = &oi
printf("%d", example)

Existe uma recomendação de deixarmos sempre a expressão de ponteiro da seguinte forma

var p * type

Ou seja, separado o asterisco, pois, evita em contextos complexos a confusão com multiplicação.
E tem um detalhe muito interessante que em GO, ao contrário de C por exemplo, não podemos realizar aritmética de ponteiros,
que causa erros tremendos de memória em C, o que torna Golang seguro.


12. if-else
	Aqui estou habituado mais que nunca, porém tem um detalhe em Go, uma completa omissão de parênteses.
Temos 3 estruturas:
	if-else
	switch-case
	select
Além disso, podemos dar um breack, ou continue. E ir além, dando um goto, para pular uma instrução.

Exemplo classico, se um numero é par ou impar

num1 := 3

    if num1 % 2 != 0 {
    	fmt.Printf("%d é impar", num1)
    } else {
    	fmt.Printf("%d é par", num1)
    }

As condicionais são do tipo boolean, então temos alguns bons exemplos do que funciona:

if true {
	// faça algo
}

if !bool1 // faça algo

if !(x == y) ou if x != y

Quando queremos retornar isso ou aquilo, podemos omitir o else:

if condição {
	return x
}
return y

Também podemos simplificar algumas coisas

if num1 := 5; num1 > 4 {
	// faça algo
}

exemplos mais completo com vários tipos:

var first int = 10
    var cond int
    if first <= 0 {
        fmt.Printf("first is less than or equal to 0\n")
    } else if first > 0 && first < 5 {
        fmt.Printf("first is between 0 and 5\n")
    } else {
        fmt.Printf("first is 5 or greater\n")
    }
    if cond = 5; cond > 10 {
        fmt.Printf("cond is greater than 10\n")
    } else {
        fmt.Printf("cond is not greater than 10\n")
    }

13. Testando erros com funções
	Aqui, temos uma estrutura que confesso não lembrar ter vista antes. Para analisarmos
um success ou err, usaremos a seguinte estrutura como exemplos:

anInt, erre = strconv.Atoi(str)

Nesse caso, se str não puder ser convertido para int, será retornado 0 para anInt e erro absorve error.

Segue um exemplo completo que, quando imprimimos a variavel err, ela tem a seguinte mensagem:

"strconv.ParseInt: parsing "ABC": invalid syntax"

codigo:
	var orig string = "ABC"
    var an int
    var err error
    // storing integer and error information
    an, err = strconv.Atoi(orig)

    if err != nil { // if it was an error, discontinue
        fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
        fmt.Println(err)
        return
    }
    fmt.Printf("The integer is %d\n", an)

Outro exemplo com vários tratamentos de erros:

import (
    "fmt"
    "errors"
)

// Factorial calculates the factorial of a non-negative integer n.
// It returns 0 as the factorial value if n is negative.
func Factorial(n int) (int, error){
    if n < 0 {
        return 0, errors.New("fatorial não pode ser negativa.")
    }

    if n == 0 {
        return 1, nil   // Factorial of 0 is 1
    }

    factorial := 1
    for i := 1; i <= n; i++ {
        factorial *= i
    }

    return factorial, nil
}

func main() {
    n := 10
    x := -10

    result, err := Factorial(n)
    if err == nil {
        fmt.Printf("Factorial of %d is %d\n", n, result)
    }
    result, err = Factorial(x)
    if err == nil {
        fmt.Printf("Factorial of %d is %d\n", x, result)
    }
}

14. Switch-case
	Aqui é da mesma forma, até o momento sem misterio.

switch var1 {
case val1:

case val2:

default:

}

Se eu quiser que seja executado outro case mesmo de valor diferente, eu posso utilizar fallthrough

E se formos atribuir algo em switch, utilizamos o pontoe virgula:

 switch num1 := 100;{
        case num1 < 0:
            fmt.Println("Number is negative")

        case num1 > 0 && num1 < 10:
            fmt.Println("Number is between 0 and 10")

        default:
            fmt.Println("Number is 10 or greater")
    }

Challenge: Season of a Month
This lesson brings you a challenge to solve.

Problem statement
Write a function Season that has as input-parameter of a month number and which returns the name of the season to which this month belongs (disregard the day in the month). Make sure you follow the following criteria of month names and their values:

svg viewer

January: 1
February: 2
March: 3
April: 4
May: 5
June: 6
July: 7
August: 8
September: 9
October: 10
November: 11
December: 12
Another thing to note is the seasons of the months. Look at the following mapping:

Winter: 1, 2, and 12
Spring: 3, 4, and 5
Summer: 6, 7, and 8
Autumn: 9, 10, and 11
Input
An int variable

Output
A string variable


Meu código solução foi:

func Season(month int) string {
  switch month{
    case 1, 2, 12:
      return "Winter"

    case 3, 4, 5:
      return "Spring"

    case 6, 7, 8:
      return "Summer"

    case 9, 10, 11:
      return "Autumn"
  }
  return "Season unknown"
}

A resposta do autor foi:

func Season(month int) string {
	switch month {  // switch on the basis of value of the months(1-12)
		case 12,1,2:  return "Winter"  // Jan, Feb and Dec have winter
		case 3,4,5:	  return "Spring"  // March, Apr and May have spring
		case 6,7,8:   return "Summer"  // June, July and Aug have summer
		case 9,10,11: return "Autumn"  // Sept, Oct and Nov have autumn

		default: return "Season unknown"         //value outside [1,12], then season is unkown
	}
}

De fato apesar de ter a mesma lógica, poderia ter melhorado visualmente o meu código.

15. For
	Provavelmente uma das mais importantes estruturas de controle na programação, o laço for.
Temos dois tipos de laço for:
	- iteração por contador
	- iteração por condicional

15.1. Iteração por contador:
	Aqui é a mais que habituado, dá até uma nostalgia rever esses assuntos como um principiante (em GO eu sou um principante, não tenho de ser
arrogante assim). TUdo consiste na seguinte estrutura:
	for inicialização; condição; modificação {}

Por exemplo:
	for i := 0; i < 5; i++ {}

E também podemos usar mais de um contador:

for i, j := 0, N; i < j; i, j = i+1, j-1 {}

15.2. Iteração por condição
	Aqui é bem mais simples, o exemplo diz por si só:

	var i int = 0

	for i < 5 {
		i = i + 5
	}

15.3. uso com intervalo, for range
	for range é uma construção iteradora em GO (por favor, se alguém ler isso, é fundamental saber o que é iteração),
uma variação interessante para criar um loop sobre cada item de uma coleção
Exemplo geral:

for ix, val := range coll {}

Explicando por outras linhas ,ele consegue percorrer facilmente variaveis com strings e imprimir o index de cada, como se fosse um array.
Codigo:

str3 := "Vinicius Aarao"
    for pos, char := range str3 {
        fmt.Printf("Minha letra %c na posição %d\n", char, pos)
    }

15.4. Break e continue
	Break encerra o loop condicional, indi para fora do loop.

for i:=0; i<3; i++ {    // outer loop
        for j:=0; j<10; j++ {   // inner loop
            if j>5 {
                break                   // breaking inner loop
            }
            fmt.Printf("%d",j)
        }
        fmt.Printf(" ")
    }

esse exemplo é perfeito, só peguei do autor discarado.

	Já continue é exatamente o oposto, mesmo que a condição seja satisfeita, ele continua dentro do corpo do loop.

for i := 0; i < 10; i++ {     // for loop
        if i == 5 {
            continue                      // continuing the loop
        }
        fmt.Printf("%d",i)
        fmt.Print(" ")
    }


16. Labels - goto
	rótulo é um sequencia de caracteres que identifica um local dentro de um código. Uma linha de código que começa
com for, switch ou select, pode ser precedida por um label no seguinte formato:

IDENTIFICADOR:

// restante do código

Um exemplo do uso de label:

LABEL1:
for i := 0; i <= 5; i++ {
	for j := 0; j <= 5; j++ {
		if j == 4 {
			continue LABEL1
		}
		fmt.Printf("i = %d\n, j = %d\n, i, j)
	}
}

Existe também o goto, nada mais é do que o go to here:

    i:=0
    HERE:           // adding a label for a location
    fmt.Printf("%d",i)
    i++
    if i==5 {
        return
    }
    goto HERE      // goto HERE

MINI-CHALLENGE

Crie um programa que simule uma transação bancária com tratamento de erros e etapas de transação claras, facilitando a transferência de fundos entre duas contas. Este programa deve aceitar o valor da transferência, o saldo da conta de origem e o saldo da conta de destino como parâmetros.

Implemente uma função chamada transferFundsque manipula corretamente o seguinte:

Verificar saldo: Verifique se a conta de origem tem saldo suficiente para a transferência. Caso contrário, trate os erros de saldo insuficiente com cuidado.
Atualizar saldo: deduza o valor da transferência do saldo da conta de origem e adicione-o ao saldo da conta de destino.
Conclusão: marque a conclusão bem-sucedida da transação.

type Account struct {
    ID      int
    Balance float64
}

func transferFunds(amount float64, source *Account, destination *Account) {
    if amount > source.Balance {
        fmt.Println("Saldo Insuficiente")
    } else {
        source.Balance -= amount
        destination.Balance += source.Balance
        fmt.Printf("Saldo atual: %f", source.Balance)
    }
}

Eu não detalhei muito mas, essa foi minha solução.
*/
