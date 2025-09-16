package Server

import (
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
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

/*
1. TCP server
	GO é muito útil para escrever aplicações web.

Esse package é um repositório de exemplo para uma aplicação simples cliente-servidor usando
o protocolo TCP e goroutine. Além disso, iremos precisar do pacote net, que contém
métodos para trabalhar com TCP/IP e UDP.

Exemplo simples:

	fmt.Println("Starting the server")
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}
		fmt.Printf("Received: %v", string(buf))
	}
}

Exemplo mais elaborado:

const maxRead = 25

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}

func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: `"+hostAndPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	println("Listening to:    ", listener.Addr().String())
	return listener
}

func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	println("Connection from: ", connFrom)
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0 // to prevent overflow
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN: // try again
			continue
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connFrom)
	checkError(err, "Close: ")
}

func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "Write: wrote "+string(wrote)+" bytes.")
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
	}
}

func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}

2. HTTP
	O GO para o propósito de http, tem o package net/http. Para algo simples como um famoso
hello world, o servidor é iniciado com http.ListenAndServe("localhost:8080", nil), vai retornar nil
se tiver tudo ok.
	Um endereço web é representado pelo tipo http.URL, que possui o Path que contém a URL string. As
solicitações de clientes são descritas pelo tipo http.Request, que possui uma URL,
um Header, um Body e outras coisas.
	Se a solicitação req for POST de um formulário HTML e var1 for o nome de um campo de entrada HTML, o valor
inserido pode ser capturado por req.FormValue("var1").

var1, found := request.Form["var1"]

Exemplo de http simples:


*/
