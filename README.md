# golang-api

Este projeto foi criado com o objetivo de ajudar a entender como funciona um CRUD básico em Go usando o framework Fiber e o ORM GORM.

## Requisitos

Antes de começar, você precisará ter o Go instalado no seu sistema. Para instalar o Go, visite [golang.org](https://golang.org/doc/install).

## Instalação

1. Clone o repositório para o seu sistema:

```bash
git clone https://github.com/KayoRonald/golang-api
```
2. Acesse o diretório do projeto:

```bash
cd golang-api
```

3. Inicie o servidor:

```bash
go run main.go`
```

## Endpoints

A API possui os seguintes endpoints:

| Método | Endpoint | Descrição |
| ------ | -------- | --------- |
| GET | /book | Retorna todos os books |
| GET | /book/:id | Retorna um book pelo ID |
| POST | /book | Cria um novo book |
| PUT | /book/:id | Atualiza um book existente |
| DELETE | /book/:id | Deleta um book pelo ID |

Os dados são retornados em formato JSON.

## Contribuição

Se você encontrar algum bug ou tiver alguma sugestão de melhoria, sinta-se à vontade para abrir uma issue ou um pull request.

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo LICENSE para obter mais informações.

Espero que isso ajude a criar um README.md para o seu projeto em Go usando o Fiber e o GORM!