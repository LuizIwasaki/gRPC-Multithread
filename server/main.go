package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	pb "service/grpc/pb"
)

type Server struct {
	pb.UnimplementedServiceHelloServer
}

func (s *Server) ValidateCPF(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
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
	listenAddr := ":0" // Porta em que o servidor gRPC será iniciado

	// 	lis, err := net.Listen("tcp", listenAddr)
	// if err != nil {
	//     log.Fatalf("Failed to listen: %v", err)
	// }

	// // Obtém a porta atribuída pelo sistema
	// port := lis.Addr().(*net.TCPAddr).Port
	// fmt.Printf("Server listening on port %d\n", port)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Obtém a porta atribuída pelo sistema
	port := lis.Addr().(*net.TCPAddr).Port
	fmt.Printf("Server listening on port %d\n", port)
}
