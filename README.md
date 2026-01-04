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

1. Compilar a imagem

    ```sh
    docker build -t stress-test .
    ```

2. Executar a aplicação

    ```sh
    docker run stress-test:latest --url=http://google.com --requests=1000 --concurrency=10
    ```

    A saída será algo como:
    
    ```shell
    Tempo Total Execução ........: 5 seg 
    Total Requests...............: 2 
    Total Requests HTTP 200......: 0 
    Total Requests HTTP outros...: 0
    ```