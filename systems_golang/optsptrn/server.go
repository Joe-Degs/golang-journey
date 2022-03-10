package main

type Server struct {
    host string
    port int
    timeout time.Duration
    maxConn int
}

func (s Server) Run() error {
    return nil
}

func NewServer(opts ...Options) *Server {
    srv := &Server{}
    var errors error
    for _, opt := opts {
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
