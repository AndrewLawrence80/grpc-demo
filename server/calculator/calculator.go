package calculator

import (
	"context"
	"fmt"

	pb "github.com/andrewlawrence80/grpc-demo/proto/calculator"
)

type CalculatorServerImpl struct {
	pb.UnimplementedCalculatorServer
}

func (CalculatorServerImpl) Add(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{Result: req.Num1 + req.Num2}, nil
}
func (CalculatorServerImpl) Subtract(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{Result: req.Num1 - req.Num2}, nil

}
func (CalculatorServerImpl) Multiply(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{Result: req.Num1 * req.Num2}, nil

}
func (CalculatorServerImpl) Divide(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	if req.Num2 == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	return &pb.CalcResponse{Result: req.Num1 / req.Num2}, nil
}
