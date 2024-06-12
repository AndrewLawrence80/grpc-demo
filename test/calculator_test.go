package test

import (
	"context"
	"testing"

	pbCalculator "github.com/andrewlawrence80/grpc-demo/proto/calculator"
	calculator "github.com/andrewlawrence80/grpc-demo/server/calculator"
)

func TestAdd(t *testing.T) {
	server := &calculator.CalculatorServer{}
	req := &pbCalculator.CalcRequest{Num1: 1, Num2: 2}
	resp, err := server.Add(context.Background(), req)
	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	expected := float32(3)
	if resp.Result != expected {
		t.Errorf("expected %v, got %v", expected, resp.Result)
	}
}
