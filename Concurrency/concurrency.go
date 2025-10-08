package Concurrency

import (
	"fmt"
	"sync"
)

/*
1. Concurrency
	Concorrência por definição, é a capacidade de decompor um programa de computador ou algoritmo
em partes individuais, que podem ser executados independentemente. Por isso, garantir a exatidão do programa
é fundamental, além da sincronização e coordenação dos processos, aumentando a eficácia e evitando um possível
deadlock em casos de dependência.
	Um exemplo conhecidíssimo de concorrência é o sistema bancário, nas situações por exemplo de deposit() e withdraw().

2. Simultaneidade e paralelismo
	"Simultaneidade é a capacidade de lidar com muitas coisas ao mesmo tempo
	e paralelismo é a capacidade de fazer muitas coisas ao mesmo tempo."

	Paralelismo ocorre quando dividimos uma tarefa em subtarefas e as executamos simultaneamente. A programação
paralela é difícil de alcançar, por isso entra a simultaneidade.
	"COncorrência é sobre estrutura
	paralelismo é sobre execução."

Vamos para um exemplo com os dois:

Na simultaneidade, é como tomar café da manhã e, ao invés de fazer 1 item de cada vez, podemos coordenar entre
ferver a agua enquanto preparo o sanduíche.

Paralelismo é como ter 2 pessoas fazendo o café da manhã, uma preparando o sanduíche e a outra fervendo a água.

3. Data Races and Race COnditions
	Uma corrida de dados acontece quando dois ou mais threads acessam um recurso compartilhado, e um dos acessos
é de escrita.
	O GO tem um  race detector, o exemplo abaixo mostrará ele:

func deposit(balance *int, amount int){
	*balance += amount
}

func withdraw(balance *int, amount int){
	*balance -= amount
}

func main (){

	balance := 100
	go deposit(&balance,20)
	withdraw(&balance,30)

	fmt.Println("Balance:", balance)

3, Deadlocks e Starvation
	Deadlock ocorre quando os processos estão esperando uns pelos outros para liberar recursos, e nenhum deles
pode continuar. Starvation é quando um processo não consegue acessar os recursos necessários para continuar sua
execução, porque outros processos estão monopolizando esses recursos.

4. Básico de concorrência com Goroutines
	Utilizamos as Goroutines para lidar com as concorrências. Um exemplo introdutório de aplicação:

func WelcomMessage(){
	fmt.Printl("welcome to Go Concurrency")
}

func main(){
	go WelcomMessage()
	fo func(){
		fmt.Println("Main function")
	}()
}

Nesse caso, não haverá saída, pois a goroutine pode não ser executada antes do término da função main.
Se acrescentermos:

time.Sleep(1000)

Poderemos ver ambas as saídas na tela. Isso ocorre porque houve um atraso de tempo que permite as goroutines terminarem antes de saírmos
da execução principal.

5. Fork-Join
	Go utiliza fork-join de simultaneidade. A função main é a goroutine principal, e quando uma nova goroutine é iniciada, ela é "forked" (criada) a partir da goroutine principal. Quando a goroutine principal termina, todas as goroutines filhas também terminam.
Isso significa que se a função main terminar antes das outras goroutines, elas serão interrompidas.

A thread/processo filho se separa da sua thread/processo pai e, após a execução, retorna para o chamado ponto de junção.

As Goroutines tem o mesmo espaço de endereço, o programador tem a função de implementar a simultaneidade, tendo em mente a prevenção
de condições de corrida e deadlocks.

Além disso, as goroutines como thread leves, são mais eficientes em termos de memória e desempenho do que as threads tradicionais, já que tem
sua própria pilha de chamadas.

6. Channels
	Channel é um channel entre goroutines para sincronizar a execução e se comunicar enviando/recebendo dados.
Os canais são baseados na ideia de Processos Sequenciais de Comunicação (CSP), proposta por Hoare em 1978.

channelName := make(chan datatype)

para enviar dados:

channelName <- data

para receber dados:

data := <- channelName

Exemplo de envio e recebimento:

func sendValues(myIntChannel chan int){
	for i := 0; i < 10; i++ {
		myIntChannel <- i
	}
}

func main(){
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for i := 0; i < 10; i++ {
		fmt.Println(<- myIntChannel)
	}
}

Nesse caso, a função sendValues envia valores para o canal <-, enquanto na função main, myIntChannel
recebe os valores do canal <-.

Um truque sintático para fechar o canal de maneira segura:

func sendValues(myIntChannel chan int){
	for i := 0; i < 10; i++ {
		myIntChannel <- i
	}
	close(myIntChannel)
}

func main(){
	myIntChannel := make(chan int)

	go sendValues(myIntChannel)

	for value := range myIntChannel {
		fmt.Println(value)
	}
}

Outra boa prática é usar defer para fechar o canal.

func main(){
	myIntChannel := make(chan int)
	defer close(myIntChannel)
	go sendValues(myIntChannel)
	for value := range myIntChannel {
		fmt.Println(value)
	}
}

E existe um canal com buffer:

channelName := make(chan int, 2)

Isso significa que pode receber até 2 valores. Acima disso, dará deadlock. Assim como envio vazio, também pode dar deadlock.

*/

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	done := make(chan bool)
	mid := len(data) / 2
	var left []int

	go func() {
		left = MergeSort(data[:mid])
		done <- true
	}()
	right := MergeSort(data[mid:])
	<-done
	return Merge(left, right)
}

func main() {
	// Example usage
	data := []int{38, 27, 43, 3, 9, 82, 10}
	sortedData := MergeSort(data)
	fmt.Println("Sorted Data:", sortedData)
}

/*
7. WaitGroups
	Uma WaitGroup espera por um conjunto de goroutines terminarem. O contador é incrementado quando uma goroutine é iniciada e decrementado quando a goroutine termina.
A função Wait bloqueia até que o contador chegue a zero.

Podemos usar com as seguintes funções:

.Add(int): função recebe um valor inteiro que é essencialmente o número de goroutines. Deve ser executada antes de uma goroutine.

.Done(int): função chamada dentro da goroutine para sinalizar que foi executada com sucesso.

.Wait(): bloqueia o programa até que todas as goroutines tenham sido executadas.

Exemplo completo:
*/

func WelcomeMessage() {
	fmt.Println("oi")
}

func mainWG() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		WelcomeMessage()
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello World")
		wg.Done()
	}()

	wg.Wait()
}

/*
8. Mutex
	Um Mutex serve para proteger seções críticas do código, impedindo que outros processos entrem em uma seção crítica de dados enquanto
 outro processo já está nela. Isso é feito através de dois métodos principais: Lock() e Unlock().
	Usando como exemplo um depósito e saque bancário, não podemos realizar ambos ao mesmo tempo, por isso, cada vez
que lidamos com o saldo, a seção de código é crítica. Em GO, se uma goroutine estiver em uma seção crítica, o mutex não permitirá que nenhuma
outra goroutine entre nessa seção até que a primeira tenha terminado e liberado o mutex.
	Exemplo:
*/

func deposit(balance *int, amount int, myMutex *sync.Mutex, myWaitGroup *sync.WaitGroup) {
	myMutex.Lock()
	*balance += amount
	myMutex.Unlock()
	myWaitGroup.Done()
}

func withdraw(balance *int, amount int, myMutex *sync.Mutex, myWaitGroup *sync.WaitGroup) {
	myMutex.Lock()
	*balance -= amount
	myMutex.Unlock()
	myWaitGroup.Done()
}

func mainMutex() {
	balance := 100
	var myWaitGroup sync.WaitGroup
	var myMutex sync.Mutex

	myWaitGroup.Add(2)
	go deposit(&balance, 20, &myMutex, &myWaitGroup)
	withdraw(&balance, 30, &myMutex, &myWaitGroup)

	myWaitGroup.Wait()
	fmt.Println("Balance:", balance)
}
