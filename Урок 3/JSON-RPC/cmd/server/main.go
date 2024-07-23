package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/moguchev/microservices_courcse/jsonrpc/model"
)

// Service - our service
type Service struct{}

// Multiply is the method that will be called via JSON-RPC.
func (s *Service) Multiply(r *http.Request, args *model.MultiplyRequest, reply *model.MultiplyResponse) error {
	// *reply = model.MultiplyResponse{
	// 	Value: args.A * args.B,
	// }

	return errors.New("error")
}

func main() {
	// Create a new RPC server
	server := rpc.NewServer()

	impl := new(Service)
	// Register the ArithService
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterService(impl, "")

	// Register the RPC handler to handle requests on the "/rpc" endpoint
	http.Handle("/rpc", server)

	log.Println("Server running on port 8080")

	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}

	// curl -X POST http://localhost:8080/rpc -H 'cache-control: no-cache' -H 'content-type: application/json' -d '{"method": "Service.Multiply","params": [{"a":5,"b":2}],"id": "1"}'
}
