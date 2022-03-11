package main

import (
    "net"
    "strconv"
)

// that exposes the error part as well
type Option func() (func(*Server), error)

func Host(host string) Option {
    return func() (func(*Server), nil) {
	   return func(s *Server) {
		  s.host = host
	   }, nil
    }
}

func Port(port uint16) Option {
    return func(s *Server) {
	   s.port = port
    }
}

func HostPort(hs string) Option {
    var err error
    return func(s *Server) {
    }, nil
}

func HostPort(host string) Option {
    return func() (func(*Server), nil) {
	   var err error
	   return func(s *Server) { 
		  h, s, e := net.SplitHostPort(hs)
		  if e != nil {
			 err = e
			 return
		  }
		  s.host = h
		  strconv.ParseUint(p, 10, 16)
		  s.port = port
	   }, nil
    }
}

func Timeout(timeout time.Duration) Option {
    return func (s *Server) {
	   s.timeout = timeout
    }
}

func MaxConn(max int) Option {
    return func(s *Server) Option {
	   s.maxConn = max
    }
}
