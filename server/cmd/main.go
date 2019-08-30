package main

import (
	"../../server"
)

func main()  {
	var s server.Server
	s = server.NewServer()
	s.Listen(":3333")
	s.Start()
}