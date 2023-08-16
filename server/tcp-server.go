package server

import (
	"log"
	"net"
	"os"

	"github.com/Sparsh1401/go-DanteDb/config"
)

func RunSyncTcpServer() {
	log.Println("Running a Tcp Server on Port:7379")
	var conn_count int = 1
	lis, err := net.Listen(config.Type, config.HOST+":"+config.PORT)

	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		conn_count += 1
		log.Println("connection established with client", conn.RemoteAddr(), "no of client", conn_count)
		handleConnection(conn, conn_count)

	}
}

func handleConnection(conn net.Conn, conn_count int) {
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer[:])
	if err != nil {
		conn.Close()
		conn_count -= 1
		log.Println("Error with the connection", conn.RemoteAddr(), "Closed the connection", conn_count)

	}

	log.Println("Message Recieved", string(buffer[:n]))

	conn.Write([]byte("Did you send: " + string(buffer[:])))

}
