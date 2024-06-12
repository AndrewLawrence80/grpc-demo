package test

import (
	"context"
	"testing"

	pbGreeter "github.com/andrewlawrence80/grpc-demo/proto/greeter"
	greeter "github.com/andrewlawrence80/grpc-demo/server/greeter"
)

func TestSayHello(t *testing.T) {
	server := &greeter.GreeterServer{}
	req := &pbGreeter.HelloRequest{Name: "World"}
	resp, err := server.SayHello(context.Background(), req)
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	expected := "Hello World"
	if resp.Greeting != expected {
		t.Errorf("expected %v, got %v", expected, resp.Greeting)
	}
}
