package main

import (
    "net"
    "strconv"
    "time"
)

// that exposes the error part as well
type Option func() (func(*Server), error)

func Host(host string) Option {
    return func() (func(*Server), error) {
	   return func(s *Server) {
		  s.host = host
	   }, nil
    }
}

func Proto(proto string) Option {
    return func() (func(*Server), error) {
	   return func(s *Server) {
		  s.proto = proto
	   }, nil
    }
}

func Port(port uint16) Option {
    return func() (func(*Server), error) {
	   return func(s *Server) {
		  s.port = port
	   }, nil
    }
}

func HostPort(host string) Option {
    return func() (func(*Server), error) {
	   var err error
	   return func(s *Server) { 
		  hs, pr, er := net.SplitHostPort(host)
		  if er != nil {
			 err = er
			 return
		  }
		  s.host = hs
		  var port uint64
		  port, err = strconv.ParseUint(pr, 10, 16)
		  if err != nil { return }
		  s.port = uint16(port)
	   }, err
    }
}

func Timeout(timeout time.Duration) Option {
    return func() (func(*Server), error) {
	   return func (s *Server) {
		  s.timeout = timeout
	   }, nil
    }
}

func MaxConn(max int) Option {
    return func() (func(*Server), error) {
	   return func(s *Server) {
		  s.maxConn = max
	   }, nil
    }
}
