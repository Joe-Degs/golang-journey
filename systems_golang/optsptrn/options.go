package main

// that exposes the error part as well
type Option func() (func(*Server), error)

func Host(host string) Option {
    return func() (func(*Server), nil) {
	   return func(s *Server) {
		  s.host = host
	   }, nil
    }
}

func Port(port int) Option {
    return func(s *Server) {
	   s.port = port
    }
}

func HostPort(hs string) Option {
    return func(s *Server) {
	   h, s, err := net.SplitHostPort(hs)
	   if err != nil {
		  return err
	   }
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
