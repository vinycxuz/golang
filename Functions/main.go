package Functions

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	example := &ListNode{1, &ListNode{1, &ListNode{3, nil}}}
	example2 := &ListNode{2, &ListNode{3, &ListNode{4, nil}}}

	mergeTwoLists(example, example2)
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	endList := new(ListNode)
	list11 := list1
	list22 := list2
	for list11 != nil {
		if list11.Val <= list22.Val {
			endList.Val = list11.Val
			endList.Next = list22
		} else {
			endList.Val = list22.Val
			endList.Next = list11
		}
	}
	imprimirList(endList)
	return endList
}

func imprimirList(list *ListNode) {
	atual := list
	for atual != nil {
		fmt.Printf("%d", atual.Val)
		atual = atual.Next
	}
	fmt.Println("nil")
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

Podemos depurar com closures também utilizando pacotes runtimes e logs, e isso se torna muito útil para saber qual função está sendo executada.
Utilizando o pacote runtime para pegar as informações.

where := func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
	}
where()

where()

vocÊ tem uma sapida semelhante a isso:
2025/08/26 00:36:27 /usercode/main.go:11
2025/08/26 00:36:27 /usercode/main.go:14
2025/08/26 00:36:27 /usercode/main.go:19

Ou utilizando .log como já foi visto anteriormente.

7. Optimizing programs
	Podemos cronometrar o tempo de execução de determinado cálculo utilizando Now() do pacote time. E então calcular a diferenã entre eles,
utilizando Sub().

func Calculation(){
	for i:= 0; i < 1000; i++ {
		fmt.Println("oi")
	}
	start := time.Now()
	Calculation()
	end := time.Now()
	delta := end.Sub(start)

8. Array
	Array não é tão utilizado quanto slice em GO, diz-se que raramente se vê. Mas vamos aprender,
a sintaxe:

var arr1 [10]int

E podemos:

arr1[2] = 99

TAmbém podemos declarar um array com new(), lembrando que ele tem zero de espaço alocado e retorna o endereço,
logo, ficaria algo assim:

var arr2 = new([5]int)

Exemplo:

func f(a [3]int) { fmt.Println(a) }   // accepts copy

func fp(a *[3]int) { fmt.Println(a) } // accepts pointer

func main() {
  var ar [3]int
  f(ar) // passes a copy of ar
  fp(&ar) // passes a pointer to ar
}

Existe algumas variantes para especificar diferentes tamanhos do array.

var arrAge = [5]int{29, 12, 44, 25, 40}

A segunda variante é quando o compilador precisará contar quantos itens tem no array:

var arrAge2 = [...]int{23, 67}

E podemos criar chave:valor padrão, muito conhecido na programação:

var arrAge3 = [...]string{3: "eu:, "4": ela}

9. Slices
	Slices são referências a um segmento de código de um array. Slices na memória é uma estrutura com 3 campos:
	- ponteiro para a matriz subjacente
	- comprimento do slice
	- capacidade do slice

Ele obviamente precisa de um array subjacente para existir. A declaração se dá por:

var identifier []type

ou

var slice1 []type = arr1[start:end]

Se eu atribuir um array completo para um slice, também posso escrever:

var slice1 []type = arr1[:]

Para pegar fatias do array com valores especificos, podemos utilizar:

s := [3]int{1,2,3}[:]

Também podemos criar slice com make()

var slice1 []type = make([]type, len)
no caso:

slice := make([]int, 10)

Utilizamos new() em tipos de valor como matrizes e estruturas enquanto make() retorn um valor inicializado. Resumindo, new aloca e make cria.

9.1. Bytes e Slices
	Slices of bytes é tão comum no GO que tem um pacote para manipular funções. Além de conter algo muito útil chamado BUffer:

Buffer pode ser criado como uma variável:

var buffer bytes.Buffer

um ponteiro:

var r *bytes.Buffer = new(bytes.Buffer)

ou uma função:

func NewBuffer (buf []byte) *Buffer

9.2. Copy and Append

Podemos copiar uma slice ou acrescentar valores:

func copy(dst, src[]T) int

10. Relacionando strings, arrays e slices

	Podemos criar um slice de bytes a partir de uma string atribuindo ou usando método copy()

c := []byte(str)

copy(dst[]byte, str)

	Assim como podemos também anexar uma string a um slice de byte

var b []byte
var str string
b = append(b, str)

10.1. Criando uma substring de uma string
	substr := str[inicio:fim]

10.2. Alterar o caracter de uma string
	Para alterar o caracter de uma string, transformamos em um array de bytes e após isso alterar aquele índice

s := "opa"
c := []byte(s)
c[1] = 'l'
str := string(c)

10.3. Operações com append

a) para atribuir um slice no outro
	a = append(a, b...)
b) para deletar um item no index i:
	a = append(a[:i], a[i+1:]...)
c) para cortar de i até j
	a = append(a[:i], a[j:]...)
d) para extender um slice a um novo tamanho j
	a = append(a, make([]T, j)...)
e) para inserir um item x no index i:
	a = append(a[:i], append([]T{x}, a[i:]...)...)
f) para inserir um novo slice de tamanho j até i:
	a = append(a[:i], append([]T, j), a[i:]...)...)
g) para retirar o elemento mais alto da pilha:
	x, a = a[len(a)-1], a[:len(a)-1]


11. Maps
	Para declarar um map, faremos:

var map1 map[keyvalue]valuetype

var map1 map[string]int

para atribuir um valor da seguinte forma

map1[key] = val1

para atribuir um valor do mapa a uma variável:

v := map1[key]

Exemplo mais robusto:

mapLit := map[string]int{"one": 1, "two": 2}   // making map & adding key-value pair
  var mapAssigned map[string]int
  mapCreated := make(map[string]float32)        // making map with make()
  mapAssigned = mapLit
  mapCreated["key1"] = 4.5      // creating key-value pair for map
  mapCreated["key2"] = 3.14159
  mapAssigned["two"] = 3        // changing value of already existing key
  fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
  fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
  fmt.Printf("Map assigned at \"two\" is: %d\n", mapAssigned["two"])
  fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

Mesmo crescendo dinamicamente, podemos colocar um valor máximo fixo para o map.

map2 := make(map[string]int, 100)

11.1. Verificar a existencia de uma key
	Para verificarmos, utilizamos um retorno bolean:

val1, isPresent = map1[key]

Para verificar a presença de uma chave e não se importar com o valor, podemos:

_, ok := map1[key]

ou com um if

if -, ok := map1[key]; ok {}


para deletar, basta usar a palavra reservada delete

delete(map1, key1)


Challenge
Problem statement
Make a map to hold together the number of days in a week (1 -> 7) with its name. Make sure to follow the following key-value pair pattern (KEY: VALUE):

1: Monday
2: Tuesday
3: Wednesday
4: Thursday
5: Friday
6: Saturday
7: Sunday
Then write a function that takes a key and returns the value associated with it. If a non-existent key is given, the function should return the message: Wrong Key.

Minha solução:

var Days = map[int]string{1:"Monday",2:"Tuesday",3:"Wednesday",4:"Thursday",5:"Friday",6:"Saturday",7:"Sunday"} // do initialization here

func findDay(n int) string {
	if _, key := Days[n]; key {
	    return Days[n]
	}
	return "Wrong Key"
}

11.2. Slice de Maps
	Podemos criar um slice para maps, e aqui vale verificar apenas como acessar aca item de cada map:

// Version A:
  items := make([]map[int]int, 5)
  for i := range items {
    items[i] = make(map[int]int, 1)
    items[i][1] = 2 // This 'item' data will not be lost on the next iteration
  }
  fmt.Printf("Version A: Value of items: %v\n", items)


  // Version B: NOT GOOD!

  items2 := make([]map[int]int, 5)
  for _, item := range items2 {
    item = make(map[int]int, 1) // item is only a copy of the slice element.
    item[1] = 2 // This 'item' will be lost on the next iteration.
  }
  fmt.Printf("Version B: Value of items: %v\n", items2)
}

11.3. Ordenando e invertendo maps
	Podemos usar o sort package para ordenar um map através de um slice, um map por si só não pode ser ordenado diretamente.
Para ordenar, devemos:
	- copiar as chaves para um slice,
	- ordenar e
	- imprimir as chaves com for-range.

package main
import (
    "fmt"
    "sort"
)

var (
    barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
    "echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo":
    43, "lima": 98}
)

func main() {
    fmt.Println("unsorted:")
    for k, v := range barVal {
        fmt.Printf("key: %v, value: %v / ", k, v)   // read random keys
    }
    keys := make([]string, len(barVal))     // storing all keys in separate slice
    i := 0
    for k := range barVal {
        keys[i] = k
        i++
    }
    sort.Strings(keys)  // sorting the keys slice
    fmt.Println()
    fmt.Println("\nsorted:")
    for _, k := range keys {
        fmt.Printf("key: %v, value: %v / ", k, barVal[k])   // reading key from keys and value from barVal
    }
}

E para inverter:

package main
import (
"fmt"
)

var (
barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43,
"lima": 98}
)

func main() {
  invMap := make(map[int]string, len(barVal)) // interchanging types of keys and values
  for k, v := range barVal {
    invMap[v] = k   // key becomes value and value becomes key
  }
  fmt.Println("inverted:")
  for k, v := range invMap {
    fmt.Printf("key: %v, value: %v / ", k, v)
  }
}

12. Struct
	O geral para struct é:
type identificador struct {
	field1 type1
	field2 type2
	...
}

tambem podemos simplificar da seguinte forma:

var exemplo Exemplo

Para declarar um ponteiro para uma struct, utilizamos new():

var t *T = new(T)

Para obter o valor de uma struct, basta usarmos ponto entre a struct e o campo:

exemplo.name = value

Exemplo maior com tudo:

type exemplo struct {
	age int
	weight float
	name string
}

func main (){
	ms := nem(Exemplo)
	ms.age = 25
	ms.weight = 73.5
	ms.name = "Eu"
}

Outra forma de atribuir é utilizando chaves:

exemplo1 := exemplo{ 73, 25, "Eu" }

Exemplo com parâmetros:

type Person struct {  // struct definition
  firstName string
  lastName string
}

func upPerson (p *Person) { // function using struct as a parameter
  p.firstName = strings.ToUpper(p.firstName)
  p.lastName = strings.ToUpper(p.lastName)
}

func main() {
  // 1- struct as a value type:
  var pers1 Person
  pers1.firstName = "Chris"
  pers1.lastName = "Woodward"
  upPerson(&pers1)
  fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

  // 2 - struct as a pointer:
  pers2 := new(Person)
  pers2.firstName = "Chris"
  pers2.lastName = "Woodward"
  (*pers2).lastName = "Woodward" // this is also valid
  upPerson(pers2)
  fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

  // 3 - struct as a literal:
  pers3 := &Person{"Chris","Woodward"}
  upPerson(pers3)
  fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}


12.1. Conceitos avançados com struct
	Alguns exemplo de estrutura de dados com estruturas:

- Lista e duplamente encadeada

type Node struct {
	data float64
	next *Node
}

type Node struct {
	left *Node
	inf float64
	right *Node
}

Se tivermos uma estrutura e quisermos definir um tamanho para ele, podemos utilizar Sizeof()

Além disso, podemos criar função de estrutura anônima, sem a necessidade de identificadores globais, extremamente util em escopo reduzido e não precisa
se exposto pelo resto do programa.

Segue exemplos com campos anônimos e com estrutura anonima:

type C struct {
	a int
	int
	string
}

func main() {

	var person struct {
		name string
		age  int
	}

	person.name = "Chris"
	person.age = 25

	anotherPerson := struct {
		name string
		age  int
	}{
		name: "Chris",
		age:  25,
	}

13. Métodos
	Semelhante ao conceito de POO, aqui método é uma função que atua sobre uma variável de um
determinado type, chamado de receiver.

Exemplo:

func (a *denseMatrix) Add (b Matrix) Matrix

A forma geral para criar um método é:

func (recv receiver_type) method _name(parameter_list) (return_value_list) { ...}

Exemplos completo:

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
 	two1.a = 12
	two1.b = 10
	fmt.Printf("The sum is %d", two1.AddThem())
	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is %d", two2.AddThem())
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + param
}

Outro exemplo simples:

package main
import "fmt"
import "strconv"
import "encoding/json"

type Rectangle struct { // struct of type Rectangle
    length, width int
}

func (r *Rectangle) Area() int {    // method calculating area of rectangle
    return r.length * r.width

}

func (r *Rectangle) Perimeter() int { // method calculating perimeter of rectangle
    return 2 * r.length + 2 * r.width
}

Informação interessante é que quando usando um metodo para alterar o valor interno de uma estrutura, não usamos return:
Exemplo

type employee struct {
    salary float32
}

func incremenetSalary(a* employee) increment(s float32) {
    a.salary += a.salary * s
}

13.1. Métodos getters setters
	Para usar o metodo setter, usamos o prefixo set() e para get, apenas o nome do campo.
Exemplo completo:

import (
"fmt"
"person"
)

func main() {
	p := person.Person{
		FirstName: "John",
		LastName:  "Doe",
	}

	p.SetFirstName("Jane")
	fmt.Println(p.FirstName())
}

13.2. Métodos de tipos embutidos
	É muito semelhante a herança de métodos:


type Engine struct {
	Start()
	Stop()
}

type Car struct {
	Engine
	Speed int
}

e faríamos o construct da seguinte forma:

func (c *Car) GoToWorkIn {
	c.start()
	c.Stop()
}

Um exemplo mais elaborado:

package main
import (
"fmt"
)

type Log struct {
    msg string
}

type Customer struct {
    Name string
    log *Log
}

func main() {
    c := new(Customer)
    c.Name = "Barack Obama"
    c.log = new(Log)
    c.log.msg = "1 - Yes we can!"
    // shorter:
    c = &Customer{"Barack Obama", &Log{"1 - Yes we can!"}}
    fmt.Println(c.log)
    c.Log().Add("2 - After me, the world will be a better place!")
    fmt.Println(c.Log())
}

func (l *Log) Add(s string) {
    l.msg += "\n" + s
}

func (l *Log) String() string {
    return l.msg
}

func (c *Customer) Log() *Log {
    return c.log
}

E com função anonima, basta alterar na estrutura e na atribuição para a variável.:

package main
import (
"fmt"
)

type Log struct {
  msg string
}

type Customer struct {
  Name string
  Log
}

func main() {
  c := &Customer{"Barack Obama", Log{"1 - Yes we can!"}}
  c.Add("2 - After me, the world will be a better place!")
  fmt.Println(c)
  }

func (l *Log) Add(s string) {
  l.msg += "\n" + s
}

func (l *Log) String() string {
  return l.msg
}

func (c *Customer) String() string {
  return c.Name + "\nLog:\n" + fmt.Sprintln(c.Log)
}

13.3. Multiplas heranças
	Para realizar multiplas heranças, basta ter a necessidade dos tipos pai na compilação:

type Phone struct {
	return "bere bere"
}

type Smartphone struct {
	Phone
}

func main () {
	c := new(Smartphone)
	fmt.Println(c.Phone())
}


Aprofundamento de OOP em Go: https://github.com/lanl/goop


*/
