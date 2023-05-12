# Primeiro CRUD usando Golang

Aprendendo Go, e implementando um CRUD seguindo arquitetura MVC.

## Estrutura

```
  src
  | -- controller   - Entrada e validação de dados
  | -- model        - Regras de negócio e objeto principal
  | -- view         - Gerenciamento de dados (publico/privado) e converters
  converters
  | -- tests          - Testes de integração da aplicação
  | -- configuration  - Pacote para conexões, arquivos de configuração.
  etc
  main.go
  .env
  .gitignore
```

<details>
  <summary>
    Notas de desenvolvimento
  </summary>

#
## Tópicos

<details>
  <summary>
    .ENV
  </summary>

  Para buscar um valor de uma variavel de ambiente é usado o pacote, os.GetEnv, que só "pega" os valores se uma biblioteca estiver instalada.\
  Biblioteca chamada [godotenv](https://github.com/joho/godotenv).\
  Para instalar basta usar no terminal; `go get github.com/joho/godotenv`\

</details>

<details>
  <summary>
    Biblioteca para API
  </summary>

  Para lidar com rotas será usado o gin gonic, framework, mais movimentado na comunidade.\
  Instalação: `go get -u github.com/gin-gonic/gin`

</details>

<details>
  <summary>
    Biblioteca para Validação de requisição
  </summary>

  Para validar o conteúdo enviado por requisição será usado o  [Validator](https://github.com/go-playground/validator)
  Instalação: `go get github.com/go-playground/validator/v10`

</details>

</details>

#

<details>
  <summary>
    Agradecimento
  </summary>

  Usando como base o projeto do canal [HunCoding](https://www.youtube.com/@huncoding)

</details>