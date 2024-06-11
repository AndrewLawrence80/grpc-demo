package calculator

import (
	"context"
	"fmt"

	pb "github.com/andrewlawrence80/grpc-demo/proto/calculator"
)

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (CalculatorServer) Add(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{Result: req.Num1 + req.Num2}, nil
}
func (CalculatorServer) Subtract(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{Result: req.Num1 - req.Num2}, nil

}
func (CalculatorServer) Multiply(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{Result: req.Num1 * req.Num2}, nil

}
func (CalculatorServer) Divide(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	if req.Num2 == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	return &pb.CalcResponse{Result: req.Num1 / req.Num2}, nil
}
