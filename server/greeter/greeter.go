package greeter

import (
	"context"
	"fmt"

	pbGreeter "github.com/andrewlawrence80/grpc-demo/proto/greeter"
)

type GreeterServerImpl struct {
	pbGreeter.UnimplementedGreeterServer
}

func (s GreeterServerImpl) SayHello(ctx context.Context, request *pbGreeter.HelloRequest) (*pbGreeter.HelloResponse, error) {
	return &pbGreeter.HelloResponse{
		Greeting: fmt.Sprintf("Hello %s", request.Name),
	}, nil
}
