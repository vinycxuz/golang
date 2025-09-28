package DataStructures

import "fmt"

// Lista e lista encadeada

type Node struct {
	value int
	prox  *Node
}

type List struct {
	head *Node
}

func (l *List) InserirInicio(value int) {
	newNode := &Node{value: value}
	newNode.prox = l.head
	l.head = newNode
}

func (l *List) InserirFinal(value int) {
	newNode := &Node{value: value}
	for l.head == nil {
		l.head = newNode
		return
	}
	atual := l.head
	for atual.prox != nil {
		atual = atual.prox
	}
	atual.prox = newNode
}

func (l *List) Busca(value int) bool {
	atual := l.head
	for atual != nil {
		if atual.value == value {
			return true
		}
		atual = atual.prox
	}
	return false
}

func (l *List) Imprimir() {
	atual := l.head
	for atual != nil {
		fmt.Println(atual.value)
		atual = atual.prox
	}
}

// implementação de lista duplamente encadeada

type Node2 struct {
	value int
	prev  *Node2
	next  *Node2
}

type List2 struct {
	head *Node2
	last *Node2
}

func (l2 *List2) InserirInicio2(value int) {
	newNode := &Node2{value: value}
	if l2.head == nil {
		l2.head = newNode
		l2.last = newNode
	} else {
		newNode.next = l2.head
		l2.head.prev = newNode
		l2.head = newNode
	}
}

func (l2 *List2) InserirFinal2(value int) {
	newNode := &Node2{value: value}
	if l2.head == nil {
		l2.head = newNode
		l2.last = newNode
	} else {
		l2.last.next = newNode
		newNode.prev = l2.last
		l2.last = newNode
	}
}

func (l2 *List2) ImprimirCrescente() {
	atual := l2.head
	for atual != nil {
		fmt.Println(atual.value)
		atual = atual.next
	}
}

func (l2 *List2) ImprimirDecrescente() {
	atual := l2.last
	for atual != nil {
		fmt.Println(atual.value)
		atual = atual.prev
	}
}

// arvores

type Node3 struct {
	value int
	left *Node3
	right *Node3
}
