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
