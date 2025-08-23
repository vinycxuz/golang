package Functions

func main() {

}

/*
1. Funções
	Separei o arquivo nesse ponto, pois, função é algo muito específico e uma das partes do coração
das linguagens de programação. Vale ressaltar que, neste ponto, ainda não sei como é o padrão de organizaão de diretorios
e packages em Golang. OU seja, estou indo passo a passo, sem pressa e evoluindo gradativamente (com uma velocidade um pouco
maior).
O basico é esse:

func something() {}

something()

Quando for retornar 1 ou mais argumentos:

func oneMore(cat1, cat2, cat3 int) (int, int, int) {}

Tambem podemos 'tipar' uma funcão da seguinte forma:

type binOp func( int , int) int

1.2. Parâmetros e retornos
	Indo mais afundo nesse assunto, os parametros são passados dentro dos parenteses, padrão:

func oi (nome name) {}

e podemos varias, por exemplo:

func multiplicar(num1, num2, num3 int) int {
	return num1 * num2 * num3
}

Também podemos utilizar um identificador em branco para descartar valores. Exemplo:
    var i1 int
    var f1 float32
    i1, _, f1 = ThreeValues() // blank identifier
    fmt.Printf("The int: %d, the float; %f\n", i1, f1)
}

func ThreeValues()(int, int, float32) {
    return 5, 6, 7.5
}

Challenge: Multiple Return Values
This lesson brings you a challenge to solve.

package main
import "fmt"
import "strconv"
import "encoding/json"

func SumProductDiff(i, j int)(s int, p int, d int) {   // named version
    s = i + j
    p = i * j
    d = i - j
    return s,p,d
}

Solução do adm

func SumProductDiff(i, j int)(s int, p int, d int) {
    s, p, d = i + j, i * j, i - j
    return
}

func main() {
        sum, prod, diff := SumProductDiff(3, 4)
        fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
    }


2. Defer and Tracing
	defer é uma palavra chave que controla o momento de execução de uma instrução ou função atpe o final da função.
Defer executa algo quando a função que a envolve retorna. Por exemplo:
  for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }

Esse laço for, retorna isso:
4 3 2 1 0

Ele permite garantir que certas tarefas de limpeza sejam executadas antes de retornarmos de uma função.
	- Fechando um fluxo de arquivo
	- desbloqueando um recurso boqueado
	- imprimir um rodapé em um relatório
	- fechar uma conexão com banco de dados

E serve pra inverter uma palavra sem ponteiros.

Também podemos fazer logs para executar ao final de funções

Exemplo:

package main

import (
    "fmt"
    "log"
    "os"
    "io/ioutil"
)

// openFile opens a file and returns a file pointer.
func openFile(filename string) (*os.File, error) {
    file, err := os.Open(filename)
    defer func(){
        log.Printf("aberto o arquivo %c", filename)
    }()
    if err != nil {
        return nil, err
    }
    return file, nil
}

// processFile processes the opened file.
func processFile(file *os.File) error {
    // Do some operations on the file.
    content, err := ioutil.ReadAll(file)
    defer func(){
        log.Println("feito as op op")
    }()
    if err != nil {
        return err
    }
    fmt.Println(string(content))
    return nil
}

// closeFile closes the file.
func closeFile(file *os.File) {
    if file == nil {
        file.Close()
    }
}

func main() {
    fileName := "example.txt"

    file, err := openFile(fileName)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }

    err = processFile(file)
    if err != nil {
        log.Fatalf("Error processing file: %v", err)
    }

    closeFile(file)
    fmt.Println("File processed successfully.")
}



3. Built-in functions
	Existem algumas funções integradas que podem ser usadas sem importar pacotes, como por exemplo:

close
len (entrega o comprimento de vários tipos)
cap (capacidade maxima de armazenamento com maps e slices)
new e make (alocar memória)

entre outras que serão estudadas mais a frente

4. Funções recursivas
	Funções recursivas são funções que chamam a si mesmas

package main
import (
	"fmt"
)

func main() {
	for i := uint64(0); i < uint64(22); i++ {
		fmt.Printf("Factorial of %d is %d\n", i, Factorial(i)) // calculating factorial of first 21 integers
	}
}

// named return variables:
func Factorial(n uint64) (fac uint64) {
	if n<=1{ 	//base case
		return 1
	}
	fac = n * Factorial(n-1)	// recursive case
	return
}

5. Funções de ordem superior e funções como parâmetro
	Você pode passar um valor de uma função para uma variavel da seguinte forma:

func inc1(x int) int { return x + 1 }

f1 := inc1

Umas das coisas que podemos fazer é passar uma função como parâmetro. A priori, me questionei como poderia usar. E foi rapidamente:
Exemplo:

func main() {
    callback(1, Add) // function passed as a parameter
}

func Add(a, b int) {
    fmt.Printf("The sum of %d and %d is: %d\n", a, b, a + b)
}

func callback(y int, f func(int, int)) {
    f(y, 2) // this becomes Add(1, 2)
}


6. Closures
	As vezes não precisamos nomear algumas funções (não sei se concordo), e podemos atribuir uma função a uma variável para não precisar nomea-la:

soma := func (a, b int) int { return a + b }



*/
