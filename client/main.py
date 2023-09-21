import grpc
from pb.service_pb2 import HelloRequest
from pb.service_pb2_grpc import ServiceHello_ServiceDesc

def send_request():
    # Endereço e porta do servidor gRPC
    server_address = 'localhost:50052'

    # Cria o canal de comunicação com o servidor gRPC
    channel = grpc.insecure_channel(server_address)

    # Cria o stub do serviço gRPC
    stub = ServiceHello_ServiceDesc.Stub(channel)

    # CPF de exemplo (ajuste conforme necessário)
    cpf = "12345678900"

    # Cria a mensagem de requisição
    request = HelloRequest(name=cpf)

    try:
        # Chama o método ValidateCPF do servidor
        response = stub.ValidateCPF(request)

        # Exibe a resposta
        print("Resposta do servidor:", response.greeting)

    except grpc.RpcError as e:
        # Trata possíveis erros de comunicação
        print("Erro de comunicação com o servidor:", str(e))

if __name__ == '__main__':
    send_request()
