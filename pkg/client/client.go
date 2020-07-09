package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connection, error := net.Dial("tcp", "localhost:3000")

	if error != nil {
		panic(error)
	}

	defer connection.Close()
	sendMessage(connection)
}

func sendMessage(connection net.Conn) {
	fmt.Print("> ")

	stdin := bufio.NewScanner(os.Stdin)
	if stdin.Scan() == false {
		fmt.Println("Ciao ciao!")
		return
	}

	_, error := connection.Write([]byte(stdin.Text()))

	if error != nil {
		panic(error)
	}

	var response = make([]byte, 4*1024)
	_, error = connection.Read(response)

	if error != nil {
		panic(error)
	}

	fmt.Printf("Server> %s \n", response)

	sendMessage(connection)
}
