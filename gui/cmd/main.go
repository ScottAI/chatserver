package main

import (
	"flag"

	"github.com/ScottAI/chatserver/client"
	"github.com/ScottAI/chatserver/gui"
	"github.com/gpmgo/gopm/modules/log"
)

func main() {
	address := flag.String("server","127.0.0.1:3333","address of server")
	flag.Parse()
	client := client.NewClient()
	err := client.Dial(*address)

	if err != nil {
		log.Fatal("Error when connect to server",err)
	}

	defer client.Close()

	go client.Start()
	gui.StartUi(client)
}
