# Fail Fast Validation

Este diretório contém exemplos de "Fail Fast Validation" em Go, uma abordagem de desenvolvimento que valida o input o mais cedo possível e retorna erros rapidamente caso haja falhas. Isso reduz o processamento desnecessário e melhora a eficiência do código.

### Exemplo:
O código `main.go` verifica se os parâmetros de entrada atendem aos requisitos mínimos antes de prosseguir. O erro é retornado imediatamente, evitando execuções desnecessárias.

### Como rodar:
```bash
go run main.go
```