package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	pb "service/grpc/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedServiceHelloServer
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	cpf := req.Name

	isValid := isValidCPF(cpf)

	var status string
	if isValid {
		status = "CPF válido"
	} else {
		status = "CPF inválido"
	}

	return &pb.HelloResponse{
		Greeting: status,
	}, nil
}

func isValidCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return false
	}

	sum := 0

	for i := 0; i < 9; i++ {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum += digit * (10 - i)
	}

	remainder := sum % 11
	digit1 := 11 - remainder
	if digit1 == 10 || digit1 == 11 {
		digit1 = 0
	}

	if digit1 != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum += digit * (11 - i)
	}

	remainder = sum % 11
	digit2 := 11 - remainder
	if digit2 == 10 || digit2 == 11 {
		digit2 = 0
	}

	return digit2 == int(cpf[10]-'0')
}

func main() {
	listenAddr := ":49892" // Porta em que o servidor gRPC será iniciado

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Obtém a porta atribuída pelo sistema
	port := lis.Addr().(*net.TCPAddr).Port
	fmt.Printf("Server listening on port %d\n", port)

	// Crie o servidor gRPC
	s := grpc.NewServer()
	pb.RegisterServiceHelloServer(s, &Server{})

	// Inicie o servidor gRPC em uma goroutine
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Aguarde um sinal para encerrar a aplicação
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down the server...")
	s.GracefulStop()
}
