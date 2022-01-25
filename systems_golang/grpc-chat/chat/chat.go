package chat

import (
	context "context"
	"log"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Recieved message body from client: %s", in.Body)
	return &Message{Body: "Hello Baby!"}, nil
}
