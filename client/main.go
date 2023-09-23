package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "service/grpc/pb"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:49892" // Endereço do servidor gRPC

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o CPF: ")
	cpfInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Erro ao ler o CPF: %v", err)
	}

	// Remove espaços em branco e caracteres de nova linha
	cpf := strings.TrimSpace(cpfInput)

	// Cria uma conexão segura com o servidor gRPC
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}
	defer conn.Close()

	// Cria um cliente gRPC
	client := pb.NewServiceHelloClient(conn)

	// Cria a requisição com o CPF
	request := &pb.HelloRequest{
		Name: cpf,
	}

	// Chama o método no servidor
	response, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("Erro ao chamar ValidateCPF: %v", err)
	}

	// Exibe a resposta do servidor
	fmt.Printf("Resposta do servidor: %s\n", response.Greeting)
}
