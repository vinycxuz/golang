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
