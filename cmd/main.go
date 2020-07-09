package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	port := "3000"
	listener, error := net.Listen("tcp", "localhost:"+port)

	if error != nil {
		panic(error)
	}

	fmt.Println("Server running at localhost:" + port)

	waitClient(listener)

}

// クライアントからの接続を待ち受け
func waitClient(listener net.Listener) {
	connection, error := listener.Accept()

	if error != nil {
		panic(error)
	}
	//goEchoを呼び出し
	go goEcho(connection)

	// 自身を呼び出すことでループ
	waitClient(listener)
}

// echoを実行
func goEcho(connection net.Conn) {
	defer connection.Close()
	echo(connection)
}

// コネクションから読み取った値を表示する
func echo(connection net.Conn) {

	var buf = make([]byte, 1024)

	_, error := connection.Read(buf)
	if error != nil {
		if error != io.EOF {
			panic(error)
		}
	}
	//buf内の文字列に対して、文字追加する
	s := "hey " + string(buf)

	fmt.Printf("Client> %s \n", s)

	_, error = connection.Write([]byte(s))
	if error != nil {
		panic(error)
	}

	echo(connection)
}
