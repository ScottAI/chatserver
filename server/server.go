package server

type Server interface {
	Listen(address string) error
	Broadcast(command interface{}) error
	Start()
	Close()
}