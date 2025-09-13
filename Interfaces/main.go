package Interfaces

/*
1. Introdução à interface
	GO não é uma linguagem OOP, mas apesar de não reconhecer conceitos de classes e heranças,
tem o conceito bem flexível dentro das interfaces. É a sua maneira de especificar o comportamento de um
objeto.
	Uma interface define um conjunto de métodos abstratos. Exemplo:

type Namer inferface {
	Method1(param_list) return_type
	Method2(param_list) return_type
}

As interfaces são nomeadas com 'er' no final. Interfaces no GO, normalmente devem ser curtas, no máximo 3 métodos.
Ao contrário da maioria das linguagens, interface é um tipo que pode ser passado a uma variável.

Exemplo com get e set:

package main
import (
	"fmt"
)

type Simpler interface {  // interface implementing functions called on Simple struct
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Set(u int) {
	p.i = u
}

func fI(it Simpler) int {     // function calling both methods through interface
	it.Set(5)
	return it.Get()
}

func main() {
	var s Simple
	fmt.Println(fI(&s))  // &s is required because Get() is defined with a receiver type pointer
}

Exemplo de incorporação de interface:

type Square struct {
    side float32
}

type Circle struct {
    radius float32
}

type Shaper interface {
    Area() float32
}

func main() {
    var areaInf Shaper
    sql := new(Square)
    sql.side = 10
    areaInf = sql

	if t, ok := areaInf.(*Shaper); ok {
		fmt.Println(t.Area())
	}
	if u, ok := areaInf.(*Circle); ok {
		fmt.Println(u.Area())
	} else {
		fmt.Println("Not a circle")
	}
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func (cr *Circle) Area() float32 {
    return cr.radius * cr.radius * 3.14
}

1.2. Switchs e interfaces
	Esse exemplo demonstra bem a seleção de tipos para usar corretamente o método para a estrutur apontada:

type Square struct {
    side float32
}

type Circle struct {
    radius float32
}

type Shaper interface {
    Area() float32
}

func main() {
    var areaIntf Shaper
    sql := new(square)
    sql.side = 30
    areaIntf = sql

	switch t := areaIntf (type) {
        case *Square:
            fmt.Println(t.Area())
        case *Circle:
            fmt.Println(t.Area())
        default:
            fmt.Println("Not a circle or square")
}

func (s *Square) Area() float32 {
    return s.side * s.side
}

func (c *Circle) Area() float32 {
    return c.radius * c.radius * 3.14
}

De verdade, isso é lindo. Código de desafio abaixo:

package main

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Set(u int) {
	p.i = u
}

type RSimple struct {
	i int
}

func (p *RSimple) Get() int {
	return p.i
}

func (p *RSimple) Set(u int) {
	p.i = u
}

func fI(it Simpler) int {
	it.Set()
	return it.Get()
}

func main() {
    var simple Simpler
    var s Simple
    var r RSimple

    switch i := simple (type) {
        case *Simple:
            fI(&s)
        case *RSimple:
            fI(&r)
        default:
            fmt.Println("not a or b")
    }
}

Era para eu ter usado a fI para definir se ia ser estrutura Simple ou RSimple. Resposta correta:

package main
import (
	"fmt"
)

type Simpler interface {  // interface declaring Simple methods
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Set(u int) {
	p.i = u
}

type RSimple struct {
	i int
	j int
}

func (p *RSimple) Get() int {
	return p.j
}

func (p *RSimple) Set(u int) {
	p.j = u
}

func fI(it Simpler) int { // switch cases to judge whether it's Simple or RSimple
	switch it.(type) {
	case *Simple:
			it.Set(5)
			return it.Get()
	case *RSimple:
			it.Set(50)
			return it.Get()
	default:
			return 99
	}
	return 0
}

func main() {
	var s Simple
	fmt.Println(fI(&s))  // &s is required because Get() is defined with the type pointer as receiver
	var r RSimple
	fmt.Println(fI(&r))
}

Próximo desafio concluído com sucesso, aonde tem comentário era aonde eu devera implementar


package main
import "fmt"

type Square struct {
	side float32
}

type Triangle struct {	// implement this struct
	side float32
	width float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface { // implement this interface
	Perimeter() float32
}
func main() {
	var periIntf PeriInterface
	var areaIntf AreaInterface
	sql := new(Square)
	trl := new(Triangle)
	sql.side = 5
	trl.side = 6
	trl.width = 3
	areaIntf = trl
	periIntf = sql

	if t, ok := areaIntf.(*Triangle); ok {
	    fmt.Println(t.Area())
	}
	if t, ok := periIntf.(*Square); ok {
	    fmt.Println(t.Perimeter())
	} else {
	    fmt.Println("It isn't nothing")
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 { // implement method called on square to calculate its perimeter
	return sq.side * 4
}

func (tr *Triangle) Area() float32 { // implement method called on triangle to calculate its area
	return tr.side * tr.width * 0.5
}

1.3. Interface vazia
	A priori pode ser que não faça sentido, mas nada como entender como exemplos, e uma interface vazia é equivalente a 'any' em outras linguagens.

type Anything interface{}

Exemplo com função Lambda:

type SpecialString string

var whatIsThis specialString = "oi"

func TYpeSwitch() {
	testFUnc := func(any interface()) {
		switch any.(type) {
		case string:
			fmt.Println("string")
		case int:
			fmt.Println("int")
		case SpecialString:
			fmt.Println("SpecialString")
		default:
			fmt.Println("anything")
		}
	}
	testFUnc("oi")
	testFUnc(1)
	testFUnc(SpecialString("oi"))
	testFUnc(1.0)
}

No caso acima, temos um novo tipo chamado specialString e atribuímos a ele uma string. Utilizamos a interface vazia para
ler qual o tipo da variável criada.

Com ele, podemos escrever por exemplo a estrutura inicial para uma árvore binária:

type Node struct {
    left *Node
    data interface{}
    right *Node
}

func NewNode(left, right *Node) *Node {
    return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
    n.data = data
}

1.4. Interfaces e dynamic typing
	Go é a unica que combina valores de interface, verificação de tipo estático, conversão dinâmica em tempo de execução e
não exige declaração explícita de que um tipo satisfaz uma interface. Tipos que implementam uma interface podem ser passados para qualquer função que receba essa interface
como argumento.
	Isso é semelhante à digitação de pato em linguagens dinâmicas como Python e Ruby. Exemplo:

type IDuck interface {
    Quack()
    Walk()
}

func DuckDance (duck IDuck) {
    for i := 1; i <= 5; i++ {
        duck.Quack()
        duck.Walk()
    }
}

type Bird struct {
 // ...
}

func (b *Bird) Quack() {
    fmt.Println("Quack!")
}

func (b *Bird) Walk() {
    fmt.Println("Walk!")
}

func main() {
    b := new(Bird)
    DuckDance(b)
}

	Os requisitos de implementação são verificados estaticamente pelo compilador, ou seja, verifica se um tipo implementa
todas as funções de um interface quando há uma atribuição de uma variável a uma variável dessa interface. Se por exemplo
tenhamos diferentes entidades representadas como tipos que precisam ser escritas com fluxos XML.

type xmlWriter interface {
    WriteXML(w io.Writer) error
}

func StreamXML(v interface{}, w io.Writer) error {
    if xw, ok := v.(xmlWriter); ok {
        return xw.WriteXML(w)
    }
    return encodeToXML(v, w)
}

Um exemplo que abrange como estruturas, interfaces e funções de ordem superior interagem entre si:

a. Várias vezes quando se tem uma struct no app, precisamos de uma collection de objetos dessa struct:

type Any interface {}

type Car struct {
	Model string
    Manufacturer string
    BuildYear int
}

type Cars []*Car

b. E podemos usar o fato de que funções de ordem superior podem ser argumentos para outras funções ao definir a
funcionalidade necessária em uma interface:

func (cs Cars) Process(f func(car *Car)) {
    for _, c := range cs {
        f(c)
    }
}

c. Criamos funções para obter subconjuntos e chamar Process() com um fechamento:

func (cs Cars) FindAll(f func(car *Car) bool) Cars {
    cars := make([]*Car, 0)
    cs.Process(func(c *Car) {
        if f(c) {
            cars = append(cars, c)
        }
	})
    return cars
}

d. funcionalidade de Map produzindo algo de cada objeto Car:

func (cs Cars) Map(f func(car *Car) Any) []Any {
    result := make([]Any, 0)
    ix := 0
    cs.Process(func(c *Car) {
        result[ix] = f(c)
        ix++
	})
    return result
}

e. E agora podemos definir consultas concretas como:

allNewBMWs := allCars.FindAll(func(car *Car) boll {
    return car.Manufacturer == "BMW" && car.BuildYear > 2010
})

f. Também podemos retornar funções baseadas em argumentos, talvez queiramos adicionar carros a collection com base
nos fabricantes por exemplo:

func MakeGroupedAppender(manufacturers []string) (func(car *Car, map[string]Cars) {
	groupedCars := make(map[string]Cars)
    for _, m := range manufacturers {
        groupedCars[m] = make([]*Car, 0)
    }

	groupedCars["Default"] = make([]*Car, 0)

    appender := func(c *Car) {
        if _, ok := groupedCars[c.Manufacturer]; ok {
            groupedCars[c.Manufacturer] = append(groupedCars[c.Manufacturer], c)
        } else {
            groupedCars["Default"] = append(groupedCars["Default"], c)
        }
	}
    return appender, groupedCars
}

manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
groupedAppender, sortedCars := MakeGroupedAppender(manufacturers)
allUngroupedCars.Process(groupedAppender)
BMWCount := len(groupedCars["BMW"])

1.6. Resumo de Orientação a Objetos em GO
	Go não possui classes, mas sim interfaces. Dentre os 3 conceitos mais importantes de OOP temos:
	- Encapsulamento (ocultação de dados): ao contrário de outras linguagens, GO possui apenas dois
	níveis de acesso, package scope (objeto conhecido somente dentro do pacote, fica em letra minuscula)
	e exported (objeto é visível fora do seu pacote porque começa com letra maiuscula. Um tipo só pode ter
	métodos definidos em seu pacote.
	- Herança: acontece via composição, incorporação de um ou mais tipos com o comportamento desejado.
	- Polimorfismo: por meio de interfaces nas quais uma variável de um tipo pode ser atribuída a uma variável
	que qualquer interface implementa.

2. Read and Write
	Escrita e leitura são absolutamente tudo nos softwares, seja em arquivos, buffers, fluxos de entrada e saída,
conexões de rede, pipes, entre outros.
O exemplo abaixo me lembrou muito C:

func main () {
	var firstName string
    fmt.Println("Digite seu nome: ")
	fmt.Scanln(&firstName)
	fmt.Println("oi %s", firstName)
}

para inicializar, basta dar o comando: 'go run main.go'. Também podemos usar o package bufio e os.

var inputReader *bufio.Reader
var input string
var err error

func main() {
  inputReader = bufio.NewReader(os.Stdin)
  fmt.Println("Please enter some input: ")

  input, err = inputReader.ReadString('\n')
  if err == nil {
    fmt.Printf("The input was: %s\n", input)
  }
}

3. Error-handling
	EM GO, não existe um mecanismo de try/catch, acredito que até pela segurança da linguagem e pelo
fato de ser overrused. Ao invés disso, tem o mecanismo de defer-panic-recover. E fica obviamente a critério
do dev organizar da melhor forma.

Um exemplo com err function:

if value, err := pack1.Func1(param1); err != nil {
    fmt.Printf("Error %s in pack1.Func1 with parameter %v", err.Error(), param1)
	return
} else {
	// process
}

Independente do erro, podemos fazer um error-type com errors.New, exemplo:

package main
import (
    "errors"
    "fmt"
)

var errNotFound error = errors.New("Not found error")

func main() {
    fmt.Printf("error: %v", errNotFound)
}
// error: Not found error

Exemplo em uma função:

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.News("Não pode ser numero negativo")
	}
}

Nessa caso:

if f, err := Sqrt(-1); err != nil {
	fmt.Printf("Error: %s\n", err)
}

Também podemos retornar uma espécie de fmt. Printf mas, ao invés disso, usamos fmt.Errorf

*/
