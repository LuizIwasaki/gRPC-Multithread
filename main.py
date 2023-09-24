import grpc
import sys
import os

# Adicione o caminho para a pasta pb no PYTHONPATH
sys.path.append(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'pb'))

import pb.service_pb2 as pb
import pb.service_pb2_grpc as pb_grpc

def validate_cpf_or_cnpj():
    choice = input('Informe o tipo de validação\n1 - CPF\n2 - CNPJ\n').strip()
    return choice == '1', input('Digite o número: ').strip()

def run():
    server_address = 'localhost:49892'  # Endereço do servidor gRPC em Go

    # Cria o canal de comunicação seguro com o servidor
    channel = grpc.insecure_channel(server_address)

    # Cria o cliente gRPC
    client = pb_grpc.serviceHelloStub(channel)

    try:
        is_cpf, number = validate_cpf_or_cnpj()

        if is_cpf:
            request = pb.CPFRequest(cpf=number)
            response = client.ValidateCPF(request)
        else:
            request = pb.CNPJRequest(cnpj=number)
            response = client.ValidateCNPJ(request)

        # Exibe a resposta do servidor
        print('Resposta do servidor:', response.valid)
    except grpc.RpcError as e:
        print('Erro ao chamar o servidor:', e.details())

if __name__ == '__main__':
    run()
