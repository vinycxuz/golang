package Server

func main() {

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
*/
