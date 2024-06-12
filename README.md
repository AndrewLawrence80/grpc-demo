# gRPC Demo

A gRPC demo project managed in go submodules

## Step by Step Guide

1. initialize root go module

    ```bash
    go mod init github.com/{username}/grpc-demo
    ```

2. create new directory `proto` as submodule
    ```bash
    mkdir -p proto/{greeter, calculator}
    ```

    initialize proto go module
    
    ```bash
    go mod init github.com/{username}/grpc-demo/proto
    ```

    get dependencies
    ```bash
    go get google.golang.org/grpc
	go get google.golang.org/protobuf
    ```

    create `.proto` file `proto/greeter/greeter.proto`
    ``` protobuf
    syntax = "proto3";

    package greeter;

    option go_package = "github.com/{username}/grpc-demo/proto/greeter";

    message HelloRequest { string name = 1; }
    message HelloResponse { string greeting = 1; }

    service Greeter { rpc SayHello(HelloRequest) returns (HelloResponse); }
    ```

    create `.proto` file `proto/calculator/calculator.proto`

    ```protobuf
    syntax = "proto3";

    package calculator;

    option go_package = "github.com/{username}/grpc-demo/proto/calculator";

    message CalcRequest {
        float num1 = 1;
        float num2 = 2;
    }
    message CalcResponse { float result = 1; }

    service Calculator {
        rpc Add(CalcRequest) returns (CalcResponse);
        rpc Subtract(CalcRequest) returns (CalcResponse);
        rpc Multiply(CalcRequest) returns (CalcResponse);
        rpc Divide(CalcRequest) returns (CalcResponse);
    }
    ```
    
    complie `.proto` files, in `proto/greeter` and `proto/calculator`
    ``` bash
    protoc --go_out=. --go-grpc_out=. greeter.proto 
    protoc --go_out=. --go-grpc_out=. calculator.proto 
    ```

    move the `.pb.go` file in out from sub directories, organize the `proto` directory as follows:

    ```bash
    proto/
    |--- calculator
    |   |--- calculator.proto
    |   |--- calculator.pb.go
    |   |--- calcualtor_grpc.pb.go
    |--- greeter
    |   |--- greeter.proto
    |   |--- greeter.pb.go
    |   |--- greeter_grpc.pb.go
    |--- go.mod
    |--- go.sum
    ```

    commit and push the `proto` directory under the root
    ```bash
    git commit -am "add proto file"
    git push origin master
    ```

    tag version <font color="red">(important)</font>
    ```bash
    git tag proto/v0.0.1
    git push origin proto/v0.0.1
    ```

3. create new directory `server` as submodule to manage implementations of generated `.pb.go` files

   ```bash
   mkdir -p server/{greeter,calculator}
   ```

   initialize server go module
    
   ```bash
   go mod init github.com/{username}/grpc-demo/server
   ```

   get dependencies <font color="red">(remember to add version tag @v0.0.1 defined above)</font>
   ```bash
   go get github.com/{username}/grpc-demo/proto@v0.0.1
   ```

   create `.go` file `server/greeter/greeter.go`
   
   ```go
   package greeter

   import (
   	"context"
   	"fmt"

   	pbGreeter "github.com/{username}/grpc-demo/proto/greeter"
   )

   type GreeterServer struct {
   	pbGreeter.UnimplementedGreeterServer
   }

   func (s GreeterServer) SayHello(ctx context.Context, request *pbGreeter.HelloRequest) (*pbGreeter.HelloResponse, error) {
   	return &pbGreeter.HelloResponse{
   		Greeting: fmt.Sprintf("Hello %s", request.Name),
   	}, nil
   }

   ```

   create `.go` file `server/calculator/calculator.go`
   ```go
   package calculator

   import (
   	"context"
   	"fmt"
   	pb "github.com/{username}/grpc-demo/proto/calculator"
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
   ```

   create `.go` file `server/main.go`
   ```go
   package main

   import (
   	"log"
   	"net"

   	pbCalculator "github.com/{username}/grpc-demo/proto/calculator"
   	pbGreeter "github.com/{username}/grpc-demo/proto/greeter"
   	calculator "github.com/{username}/grpc-demo/server/calculator"
   	greeter "github.com/{username}/grpc-demo/server/greeter"
   	"google.golang.org/grpc"
   	"google.golang.org/grpc/reflection"
   )

   func main() {
   	lis, err := net.Listen("tcp", ":50051")

   	if err != nil {
   		log.Fatalf("failed to listen: %v", err)
   	}

   	s := grpc.NewServer()

   	greeterServer := greeter.GreeterServer{}
   	calculatorServer := calculator.CalculatorServer{}

   	pbGreeter.RegisterGreeterServer(s, greeterServer)
   	pbCalculator.RegisterCalculatorServer(s, calculatorServer)

   	reflection.Register(s)

   	log.Println("Server is running on port :50051")
   	if err := s.Serve(lis); err != nil {
   		log.Fatalf("failed to serve: %v", err)
   	}
   }

   ```

   now the `server` directory should look like as follows:

   ```bash
   server/
   |--- calculator
   |   |--- calculator.go
   |--- greeter
   |   |--- greeter.go
   |--- main.go
   |--- go.mod
   |--- go.sum
   ```

   use debug or under `server` directory run the following command. You will see generated server is ready to run
   ```bash
   go build .
   ```
   
   commit and push the `server` directory under the root
    ```bash
    git commit -am "add server implementation"
    git push origin master
    ```

    tag version <font color="red">(important)</font>
    ```bash
    git tag server/v0.0.1
    git push origin server/v0.0.1
    ```

4. create new directory `test` as submodules to manage test files
   ```bash
   mkdir test
   ```

   initialize test go module
    
   ```bash
   go mod init github.com/{username}/grpc-demo/test
   ```

   get dependencies <font color="red">(remember to add version tag @v0.0.1 defined above)</font>
   ```bash
   go get github.com/{username}/grpc-demo/server@v0.0.1
   ```

   create `.go` file `test/greeter_test.go`
   
   ```go
   package test
   import (
   	"context"
   	"testing"

   	pbGreeter "github.com/{username}/grpc-demo/proto/greeter"
   	greeter "github.com/{username}/grpc-demo/server/greeter"
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
   ```

   create `.go` file `test/calculator_test.go`
   ```go
   package test

   import (
   	"context"
   	"testing"

   	pbCalculator "github.com/{username}/grpc-demo/proto/calculator"
   	calculator "github.com/{username}/grpc-demo/server/calculator"
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

   ```

   now the `test` directory should look like as follows:

   ```bash
   test/
   |--- calculator_test.go
   |--- greeter_test.go
   |--- go.mod
   |--- go.sum
   ```
   
   under `test` directory run the following command. You will see test result
   ```bash
   go test -v
   ```
   
   commit and push the `test` directory under the root
    ```bash
    git commit -am "add test cases"
    git push origin master
    ```

    tag version <font color="red">(important)</font>
    ```bash
    git tag test/v0.0.1
    git push origin test/v0.0.1
    ```

5. create `client` directory as submodule to manage client implementation

    ```bash
   mkdir client
   ```

   initialize client go module
    
   ```bash
   go mod init github.com/{username}/grpc-demo/client
   ```

   get dependencies <font color="red">(remember to add version tag @v0.0.1 defined above)</font>
   ```bash
   go get github.com/{username}/grpc-demo/server@v0.0.1
   ```

   create `.go` file `client/calculator_client.go`
   
   ```go
   package main

   import (
   	"context"
   	"log"
   	"time"

   	pbCalculator "github.com/{username}/grpc-demo/proto/calculator"
   	"google.golang.org/grpc"
   )

   func Calculate() {
   	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
   	if err != nil {
   		log.Fatalf("did not connect: %v", err)
   	}
   	defer conn.Close()

   	client := pbCalculator.NewCalculatorClient(conn)

   	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
   	defer cancel()

   	// Example for Add
   	addResp, err := client.Add(ctx, &pbCalculator.CalcRequest{Num1: 1, Num2: 2})
   	if err != nil {
   		log.Fatalf("could not add: %v", err)
   	}
   	log.Printf("Add Result: %f", addResp.GetResult())

   	// Example for Subtract
   	subtractResp, err := client.Subtract(ctx, &pbCalculator.CalcRequest{Num1: 5, Num2: 3})
   	if err != nil {
   		log.Fatalf("could not subtract: %v", err)
   	}
   	log.Printf("Subtract Result: %f", subtractResp.GetResult())

   	// Example for Multiply
   	multiplyResp, err := client.Multiply(ctx, &pbCalculator.CalcRequest{Num1: 3, Num2: 4})
   	if err != nil {
   		log.Fatalf("could not multiply: %v", err)
   	}
   	log.Printf("Multiply Result: %f", multiplyResp.GetResult())

   	// Example for Divide
   	divideResp, err := client.Divide(ctx, &pbCalculator.CalcRequest{Num1: 10, Num2: 2})
   	if err != nil {
   		log.Fatalf("could not divide: %v", err)
   	}
   	log.Printf("Divide Result: %f", divideResp.GetResult())
   }
   ```

   create `.go` file `client/greeter_client.go`
   ```go
   package main

   import (
   	"context"
   	"log"
   	"time"

   	pbGreeter "github.com/{username}/grpc-demo/proto/greeter"
   	"google.golang.org/grpc"
   )

   func Greet() {
   	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
   	if err != nil {
   		log.Fatalf("did not connect: %v", err)
   	}
   	defer conn.Close()

   	client := pbGreeter.NewGreeterClient(conn)

   	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
   	defer cancel()

   	r, err := client.SayHello(ctx, &pbGreeter.HelloRequest{Name: "World"})
   	if err != nil {
   		log.Fatalf("could not greet: %v", err)
   	}
   	log.Printf("Greeting: %s", r.GetGreeting())
   }
   ```

   create `.go` file `client/main.go`
   ```go
   package main

   func main() {
     Greet()
     Calculate()
   }
   ```

   now the `client` directory should look like as follows:

   ```bash
   client/
   |--- calculator_client.go
   |--- greeter_client.go
   |--- main.go
   |--- go.mod
   |--- go.sum
   ```

   <font color="red">run server</font>, under `client` directory run the following command. You will see the client will send reqeust to server and get corresponding response
   ```bash
   go build . && ./client
   ```
   
   commit and push the `client` directory under the root
    ```bash
    git commit -am "add client implementation"
    git push origin master
    ```

    tag version <font color="red">(important)</font>
    ```bash
    git tag client/v0.0.1
    git push origin client/v0.0.1
    ```