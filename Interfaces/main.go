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

1.3.

2. Read and Write
	Escrita e leitura são absolutamente tudo nos softwares, seja em arquivos, buffers, fluxos de entrada e saída,
conexões de rede, pipes, entre outros.
	GO tem uma interface para isso, chamada io.Reader e io.Writer.
*/
