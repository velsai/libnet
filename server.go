package libnet

import (
	"fmt"
	"net"
)

type Server struct {
	l    *net.TCPListener
	addr string
}

func (server *Server) listien() {
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Print(err)
	}
	server.l = ln
}

func (server *Server) AcceptTCP() {

	for {
		if tcpconn, err = server.l.AcceptTCP(); err != nil {

		}
		go servertcp(tcpconn)

	}
}
func servertcp(tcpconn *net.TCPConn) {

	fmt.Println(tcpconn)
}

func (server *Server) Loop() {

	server.listien()
	server.AcceptTCP()

}
