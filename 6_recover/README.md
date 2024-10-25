# Panic e Recover no Go

## Conceito
Em Go, `panic` e `recover` são mecanismos para lidar com erros críticos:
- **panic**: Serve para indicar um erro inesperado e interromper a execução normal do programa.
- **recover**: Captura e trata um `panic` dentro de uma função `defer`, permitindo que o programa recupere e continue a execução.

## Uso de Panic
O `panic` pode ser usado para sinalizar erros críticos, como quando ocorre um valor inválido ou um estado incorreto. Por exemplo:
```go
panic("O ano informado é inválido!")
```