package main

import (
    "net"
    "strconv"
    "time"
    "fmt"
    "strings"
    "log"
)

type Server struct {
    proto string
    host string
    port uint16
    timeout time.Duration
    maxConn int
}

func (s *Server) defProto() { s.proto = "tcp" }

func (s Server) address() string {
    if strings.Contains("tcp tcp4 tcp6 udp udp4 udp6", s.proto) {
	   return net.JoinHostPort(s.host, strconv.FormatUint(uint64(s.port), 10))
    }
    return s.host
}

func (s Server) Run() error {
    if s.proto == "" { s.defProto() }

    if strings.Contains("tcp tcp4 tcp6 unix unixpacket", s.proto) {
	   listener, err := net.Listen(s.proto, s.address())
	   if err != nil { return err }
	   switch s.proto {
	   case "tcp", "tcp4", "tcp6":
		  // set tcp socket options
		  tcpLstnr:= listener.(*net.TCPListener)
	   case "unix", "unixpacket":
		  // set unix socket options
		  unixLstnr := listener.(*net.UnixListener)
	   }
	   defer listener.Close()

	   log.Println("echo server listening on", listener.Addr().String())
	   for {
		  conn, err := listener.Accept()
		  if err != nil { log.Println(err); continue }
		  go func() {
			 return
		  }()
	   }
    } else if strings.Contains("udp udp4 udp6", s.proto) {
	   addr, err := net.ResolveUDPAddr(s.proto, s.address())
	   if err != nil { return err }
	   _, err = net.ListenUDP(s.proto, addr)
	   if err != nil { return err }
	   // set socket options and run server
    }
    return nil
}

func NewServer(opts ...Option) (*Server, error) {
    srv := &Server{}
    var errors error
    for _, opt := range opts {
	   if opt == nil { continue; }
	   optfunc, err := opt()
	   if err != nil {
		  // wrap error and continue
		  if errors != nil {
			 errors = fmt.Errorf("%w %w", errors, err)
		  } else {	
			 errors = fmt.Errorf("%w ", err)
		  }
		  continue
	   }
	   optfunc(srv)
    }
    return srv, errors
}
