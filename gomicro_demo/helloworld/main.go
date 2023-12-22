package main

import (
	"context"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type Helloworld struct{}

func (h *Helloworld) Greeting(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Message = "Hello " + req.Name + " guys..."
	return nil
}

func main() {
	consulReg := consul.NewRegistry()

	// create a new service
	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Handle(new(Helloworld)),
		// set address
		micro.Address(":8080"),
		micro.Registry(consulReg),
	)

	// initialise flags
	service.Init()

	// start the service
	service.Run()
}
