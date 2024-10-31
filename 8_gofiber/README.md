# GoFiber API Documentation

## Índice
1. [Configuração Básica do GoFiber](#configuração-básica-do-gofiber)
2. [Handlers e JSON Response](#handlers-e-json-response)
3. [Middleware CORS](#middleware-cors)
4. [Locals](#locals)
5. [Middleware Customizado](#middleware-customizado)

---

### Configuração Básica do GoFiber
Para iniciar com o GoFiber, instale a biblioteca e configure a aplicação básica:
```go
package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    app.Listen(":8080")
}
```

---

### Handlers e JSON Response
Para enviar respostas JSON e tratar requisições, crie handlers específicos para cada rota:
```go
app.Get("/api/users/user", func(c *fiber.Ctx) error {
    user := User{ID: 1, Name: "John Doe"}
    return c.JSON(user)
})
```

---

### Middleware CORS
Para permitir o compartilhamento de recursos entre origens diferentes, configure o middleware CORS:
```go
import "github.com/gofiber/cors/v2"

app.Use(cors.New(cors.Config{
    AllowOrigins: "https://www.example.com",
    AllowHeaders: "Origin, Content-Type, Accept, Authorization",
}))
```

---

### Locals
O `Locals` armazena dados temporários específicos da requisição, úteis para passar informações entre middlewares e handlers:
```go
app.Use(func(c *fiber.Ctx) error {
    c.Locals("userID", 123)
    return c.Next()
})
```

Para acessar:
```go
userID := c.Locals("userID").(int)
```

---

### Middleware Customizado
Para implementar middlewares personalizados, utilize `c.Next()` para permitir a execução da próxima função no fluxo de requisição:
```go
app.Use(func(c *fiber.Ctx) error {
    fmt.Println("Executando middleware customizado")
    return c.Next()
})
```

---

Referência de [GoFiber](https://gofiber.io/) para detalhes adicionais e exemplos avançados.
