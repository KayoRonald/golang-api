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

<div align="center">
  <img src="https://www.pragimtech.com/wp-content/uploads/2020/08/golang.png" width="190" title="Golang"/>
  <img src="https://ucarecdn.com/42dad822-770d-478e-be99-26583e67cc9c/-/format/auto/-/progressive/yes/-/preview/480x480/" width="100" alt="Fiber" title="Fiber" />
</div>

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo LICENSE para obter mais informações.