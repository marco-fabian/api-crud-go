# API CRUD em Go

Uma API simples para operações CRUD (Create, Read, Update, Delete) em posts, escrita em Go com o framework [Gin](https://github.com/gin-gonic/gin) e utilizando o banco de dados PostgreSQL.

## Índice

- [Visão Geral](#visão-geral)
- [Funcionalidades](#funcionalidades)
- [Pré-requisitos](#pré-requisitos)
- [Instalação e Execução](#instalação-e-execução)
- [Endpoints](#endpoints)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Contribuição](#contribuição)

## Visão Geral

Este projeto é uma API básica em Go que permite realizar operações CRUD em posts. Ele utiliza o framework Gin para gerenciar rotas e requisições HTTP, e se conecta a um banco de dados PostgreSQL para armazenar os dados.

## Funcionalidades

- **Criação de posts**: Crie novos posts com título, conteúdo, autor e nome de usuário.
- **Listagem de posts**: Liste todos os posts cadastrados.
- **Atualização de posts**: Atualize os dados de um post existente.
- **Exclusão de posts**: Exclua um post pelo seu ID.

## Pré-requisitos

- Go (versão 1.20 ou superior)
- PostgreSQL (ou outro banco de dados suportado)
- Variável de ambiente `DATABASE_URL` configurada com a string de conexão do banco de dados.

## Instalação e Execução

1. Clone o repositório:

   ```bash
   git clone https://github.com/marco-fabian/api-crud-go.git
   cd api-crud-go

2. Instale as dependências:

    ```bash
    go mod tidy

3. O `docker-compose.yml` já configura a variável de ambiente `DATABASE_URL` automaticamente para o serviço api.

- Basta rodar:

    ```bash
    docker-compose build
    docker-compose up -d

A API estará disponível em http://localhost:3000.

## Endpoints

- **POST** /posts: Cria um novo post.
- **GET** /posts: Lista todos os posts.
- **PATCH** /posts: Atualiza um post existente.
- **DELETE** /posts: Exclui um post pelo seu ID.

## Portas Utilizadas

- **API**: A API estará disponível em `http://localhost:3000`.
- **Banco de Dados (PostgreSQL)**: O banco de dados estará acessível em `localhost:5432`.
- **Interface do Banco de Dados (Pgweb)**: A interface gráfica para o banco de dados estará disponível em `http://localhost:3001`.

### Detalhes das Portas:

| Serviço  | Porta no Host | Porta no Container | Descrição                                   |
|----------|---------------|--------------------|--------------------------------------------|
| `api`    | 3000          | 3000               | API principal do projeto.                  |
| `db`     | 5432          | 5432               | Banco de dados PostgreSQL.                 |
| `db_ui`  | 3001          | 8081               | Interface gráfica (Pgweb) para o PostgreSQL. |

## Estrutura do Projeto

* .
    * ├── main.go
    * ├── go.mod
    * ├── go.sum
    * └── internal
        * ├── database
            * └── connection.go
        * └── post
            * ├── repository.go
            * └── service.go
        * └── internal.go

 ## Contribuição

1. Crie uma branch:

    ```bash
    git checkout -b feature/new-feature

2. Commit suas alterações:

    ```bash
    git commit -m '[feat] adding new feature'

3. Faça push para a branch:

    ```bash
    git push origin feature/new-feature

4. Abra um Pull Request.
