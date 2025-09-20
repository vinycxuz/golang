package Server

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

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.google.com/",
    "http://golang.org/"
}

func main() {
    for _, url := range urls {
        resp, err := http.Head(url
        if err != nil {
            fmt.Println(err)
	}
	fmt.Println(url, ":", resp.Status)
	}
}

Neste exemplo eu executo as requisições HEAD HTTP e retorno o status de cada um.
Algumas funções úteis que são utilizadas do pacote http:

http.Redirect(w, ResponseWriter, r *Resquest, url string, code int)

http.NotFound(w ResponseWriter, r *Request)

http.Error(w ResponseWriter, msg string, code int)

Um simples exemplo com duas chamadas de roteamento:

http.HandleFunc("/test1", SimpleServer)
http.HandleFUnc("/test2", FormServer)

package main
import (
"net/http"
"io"
)

const form = `<html><body><form action="#" method="post" name="bar">
<input type="text" name="in"/>
<input type="submit" value="Submit"/>
</form></body></html>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	switch request.Method {
		case "GET":
			io.WriteString(w, form);
		case "POST":
			io.WriteString(w, request.FormValue("in"))
}

func main() {
  http.HandleFunc("/test1", SimpleServer)
  http.HandleFunc("/test2", FormServer)
  if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
    panic(err)
  }
}

*/
