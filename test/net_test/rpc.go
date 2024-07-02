package main

import (
	"net"
	"net/rpc"
)

type UserService struct {
}

func (s *UserService) Say(request string, reply *string) error {
	*reply = "hi," + request
	return nil
}

func main() {
	var err error
	err = rpc.RegisterName("UserService", new(UserService))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":8011")
	if err != nil {
		panic(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	rpc.ServeConn(conn)
}
