package Goroutines

func main() {

}

/*
1. GOROUTINE
	Goroutine é preferível estudar separado afundo, mas, basicamente aplicamos paralelismo com 1 thread.
Go não paraleliza por padrão, sempre utiliza apenas um único núcleo e, caso queira utilizar mais, precisamos
utilizar o runtime.GOMAXPROCS()

package main
import (
    "fmt"
    "time"
)

func main() {
	fmt.Println("In main()")
    go longWait()
    go shortWait()
    fmt.Println("About to sleep in main()")
    time.Sleep(10 * le9)
}

func longWait() {
    fmt.Println("In longWait()")
    time.Sleep(5 * le9)
    fmt.Println("Exiting longWait()")
}

func shortWait() {
    fmt.println("In shortWait()")
    time.Sleep(2 * le9)
    fmt.Println("Exiting shortWait()")
}

2. Channels
	Para aumentar a utilidade das goroutines, Go possui um tipo especial chamado channel.
Channel funciona como um conduíte na qual os valores são compartilhados entre as goroutines.
Uma boa analogia é como uma esteira rolante em uma fábrica. A máquina (goroutine do produtor) coloca os itens na esteira e outra máquina
(goroutine do consumidor) pega os itens da esteira.

var identifier chan datatype

Ele tem a estrutura FIFO.

Como ele é tipo de referência (ponteiro), precisamos usar make() para alocar memória para ele.

var ch1 chan string
ch1 = make(chan string)

ch1 := make(chan string)

Para construir um canal de canais:

chanOfChans := make(chan chan string)

Quando terminamos de usar um canal, é uma boa prática fecha-lo:

close(ch1)

Além disso, tem algo bem intuitivo sobre enviar ou receber dados:

if <- ch != 1000 {
	fmt.Println("Erro")
}

Um exemplo completo:

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
    go sendData(ch)
    go getData(ch)
    time.Sleep(le9)
}

func sendData(ch chan string) {
	ch <- "Hello"
	ch <- "Hi"
    ch <- "Oi"
    ch <- "Hola"
}

func getData(ch chan string) {
    var input string
    for {
        input = <- ch
        fmt.Println(input)
        if input == "Hello" {
            break
        }
    }
    fmt.Println("Fim")
}

3. Comunicação síncrona e assíncrona
	Por padrão, a comunicação é síncrona e sem buffer, ou seja, não será concluída a operação até que tenha
recebido a resposta.

*/
