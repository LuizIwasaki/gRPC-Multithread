import grpc
import sys
import os

# Adicione o caminho para a pasta pb no PYTHONPATH
sys.path.append(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'pb'))

import pb.service_pb2 as pb
import pb.service_pb2_grpc as pb_grpc

def run():
    server_address = 'localhost:49892'  # Endereço do servidor gRPC em Go
    cpf = input('Digite o CPF: ').strip()
    # Cria o canal de comunicação seguro com o servidor
    channel = grpc.insecure_channel(server_address)

    # Cria o cliente gRPC
    client = pb_grpc.serviceHelloStub(channel)

    try:
        request = pb.HelloRequest(name=cpf)

        # Chama o método no servidor
        response = client.SayHello(request)

        # Exibe a resposta do servidor
        print('Resposta do servidor:', response.greeting)
    except grpc.RpcError as e:
        print('Erro ao chamar o servidor:', e.details())

if __name__ == '__main__':
    run()
