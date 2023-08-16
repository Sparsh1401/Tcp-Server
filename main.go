package main

import (
	"flag"
	"fmt"

	"github.com/Sparsh1401/go-DanteDb/config"
	"github.com/Sparsh1401/go-DanteDb/server"
)

func setUpFlags() {
	flag.StringVar(&config.HOST, "host", "0.0.0.0", "for DanteDb")
	flag.StringVar(&config.PORT, "Port", "7379", "for Dante Db")
	flag.StringVar(&config.Type, "type", "tcp", "server connection type")
}

func main() {
	setUpFlags()
	fmt.Println("Please wait loading your DanteDb server")
	server.RunSyncTcpServer()
}
