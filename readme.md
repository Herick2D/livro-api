# Livro API

Este projeto apresenta uma API simples para o registro e manipulação de informações de um livro. Desenvolvida em Go utilizando as tecnologias Gin, GORM, Docker e PostgreSQL. O objetivo da aplicação é realizar operações CRUD (Criar, Ler, Atualizar, Deletar) sobre dados de livros em um banco de dados.

## Considerações
Foram utilizados bibliotecas e frameworks por sua popularidade e apoio da comunidade, decisão essa tomada para fins de manutenibilidade no desenvolvimento dessa aplicação. Optei por um banco de dados relacional devido à minha familiaridade com sua estrutura e ao vasto material de suporte disponível para o desenvolvimento de APIs que utilizam bancos relacionais e seus ORMs.

## Funcionalidades

A API oferece as seguintes funcionalidades:

### 1. Criar um livro
-  **URL**: `POST /api/livros`
- **Descrição**: Cria um novo livro.
- **Requisição**:
    ```json
    {
    "titulo": "Mathematical Foundations",
    "categoria": "Física Quântica",
    "autor": "John von Neumann",
    "sinopse": "Este livro oferece uma análise rigorosa dos conceitos matemáticos essenciais para a compreensão da mecânica quântica."
    }
    ```
  
### 2. Buscar livro por ID
-  **URL**: `GET /api/livros/:id`
- **Descrição**: Retorna um livro específico.
- **Resposta de sucesso**:
    ```json
    {
    "ID": 1,
    "CreatedAt": "2025-01-06T16:49:44.496008-03:00",
    "UpdatedAt": "2025-01-06T16:49:44.496008-03:00",
    "DeletedAt": null,
    "titulo": "Mathematical Foundations",
    "categoria": "Física Quântica",
    "autor": "John von Neumann",
    "sinopse": "Este livro oferece uma análise rigorosa dos conceitos matemáticos essenciais para a compreensão da mecânica quântica."
    }
    ```
  
### 3. Atualizar um livro
-  **URL**: `PUT /api/livros/:id`
- **Descrição**: Atualiza um livro existente.
- **Requisição**:
    ```json
    {
    "titulo": "The Pragmatic Programmer: From Journeyman to Master",
    "categoria": "Fundamentos da Programação",
    "autor": "Andy Hunt",
    "sinopse": "Este livro é um guia atemporal que aborda as práticas essenciais para o desenvolvimento de software eficaz."
    }
    ```

- **Resposta de sucesso**: 
    ```json
    {
    "ID": 1,
    "CreatedAt": "2025-01-06T16:49:44.496008-03:00",
    "UpdatedAt": "2025-01-06T16:49:44.496008-03:00",
    "DeletedAt": null,
    "titulo": "The Pragmatic Programmer: From Journeyman to Master",
    "categoria": "Fundamentos da Programação",
    "autor": "Andy Hunt",
    "sinopse": "Este livro é um guia atemporal que aborda as práticas essenciais para o desenvolvimento de software eficaz."
    }
    ```
  
### 4. Deletar um livro
-  **URL**: `DELETE /api/livros/:id`
- **Descrição**: Deleta um livro existente.

### 5. Listar todos os livros
-  **URL**: `GET /api/livros`
- **Descrição**: Retorna uma lista com todos os livros cadastrados.

## Executando a aplicação

_Atenção: Garanta que você tenha o Docker instalado em sua máquina._

- Acessando o [GitHub do projeto](https://github.com/Herick2D/livro-api.git), faça o clone do mesmo;
- Acesse a pasta do projeto e execute o comando `docker-compose up --build` para criar a imagem do container e subir a aplicação;

## Observações

Esta aplicação foi desenvolvida em uma ambiente com prazos estipulados, o que pode ter limitado o tempo para ajustes detalhados e com isso pode haver comportamentos não previstos ou bugs. Caso encontre algum problema, por favor, abra uma issue no repositório do projeto.
