package main

import (
	"github.com/ScottAI/chatserver/server"
)

func main()  {
	var s server.Server
	s = server.NewServer()
	s.Listen(":3333")
	s.Start()
}