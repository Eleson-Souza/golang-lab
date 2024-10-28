# Go Routines

## Introdução
As goroutines permitem a execução simultânea de tarefas, tornando a aplicação mais eficiente ao lidar com múltiplas operações ao mesmo tempo. Elas são leves e mais baratas do que as threads tradicionais, permitindo que uma aplicação escale com menos consumo de recursos.

## Conceitos Básicos
- Uma **goroutine** é uma função que roda simultaneamente com o restante do código.
- Para iniciar uma goroutine, usa-se `go` seguido da função a ser executada.

## Controle e Sincronização
Para gerenciar e sincronizar múltiplas goroutines, usamos `sync.WaitGroup`:
- **Add(n int)**: Este método é usado para informar ao WaitGroup quantas goroutines você espera que sejam executadas. O valor n deve ser igual ao número de goroutines que você vai iniciar. Cada vez que você chama Add(), o contador interno é incrementado, sinalizando que mais tarefas estão em andamento.
- **Done()**: Chame este método dentro de uma goroutine quando ela terminar sua execução. Ele decrementar o contador do WaitGroup. Isso indica que uma das goroutines esperadas foi concluída.
- **Wait()**: Este método bloqueia a execução do programa até que o contador do WaitGroup chegue a zero, ou seja, até que todas as goroutines tenham chamado Done(). Ele é essencial para garantir que o programa não finalize antes que todas as tarefas em paralelo sejam concluídas.

## Exemplo
```go
func main() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
      defer wg.Done()
      fmt.Println("Executando tarefa 1 em goroutine")
      time.Sleep(5 * time.Second)
      fmt.Println("Tarefa 1 finalizada")
    }()

    wg.Add(1)
    go func() {
      defer wg.Done()
      fmt.Println("Executando tarefa 2 em goroutine")
      time.Sleep(5 * time.Second)
      fmt.Println("Tarefa 2 finalizada")
    }()

    wg.Wait()
    fmt.Println("Todas tarefas concluídas.")
}
```