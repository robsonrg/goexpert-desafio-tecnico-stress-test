# Desafio Pós Go Expert - Stress test

> Este projeto contém a solução para o desafio técnico sobre stress test da pós-graduação Go Expert da FullCycle.


(...) Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI

- `--url` URL do serviço a ser testado.
- `--requests` Número total de requests.
- `--concurrency` Número de chamadas simultâneas.

Execução do teste

- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.

Geração de relatório

Apresentar um relatório ao final dos testes contendo:

- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:

```sh
docker run <sua imagem docker> --url=http://google.com --requests=1000 --concurrency=10
```

# Executando a aplicação

### Pré-requisitos

- Go (>= 1.25)
- Docker (opcional)

### Instalando a aplicação localmente

1. Clone o repositório

    ```sh
    git clone <git@github.com:robsonrg/goexpert-desafio-tecnico-stress-test.git>
    
    cd goexpert-desafio-tecnico-stress-test
    ```

### Rodando localmente

1. Compile o projeto

    ```sh
    go mod download
    go build -o stresstest ./main.go
    ```

2. Dê permissão de execução

    ```sh
    chmod +x ./stresstest
    ```

2. Execute a aplicação que foi compilada
    
    ```sh
    ./stresstest --url=<$protocol://$host> --requests=<int> --concurrency=<int>
    ```

    Flags disponíveis

    | Flag          | Descrição                      |
    |---------------|--------------------------------|
    | --url         | URL do serviço a ser testado   |
    | --requests    | Número total de requisições    |
    | --concurrency | Número de chamadas simultâneas |

    Exemplo
    ```sh
    ./stresstest --url=https://google.com --requests=10 --concurrency=5
    ```

Você também poderá executar localmente, sem compilar o projeto:

```sh
go run main.go --url=https://google.com --requests=10 --concurrency=5
```

### Rodando localmente com Docker

1. Compile a imagem

    ```sh
    docker build -t stresstest .
    ```

2. Execute a aplicação

    ```sh
    docker run --rm stresstest:latest --url=http://google.com --requests=100 --concurrency=10
    ```

### Gerando relatório

A saída da execução da aplicação será um relatório assim:
    
```shell
Tempo Total Execução ........: 7 seg 
Total Requests...............: 100 
Total Requests HTTP 200......: 100 
Total Requests HTTP outros...: 0
```