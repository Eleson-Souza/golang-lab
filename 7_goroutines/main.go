package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Iniciando", name)
	time.Sleep(3 * time.Second)
	fmt.Println("Terminando", name)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go task(fmt.Sprintf("Tarefa %d", i), &wg)
	}

	wg.Wait()
	fmt.Println("\nTodas as tarefas foram concluídas")
}
