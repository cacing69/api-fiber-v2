package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Membuat instance Fiber
    app := fiber.New()

    // Menangani rute GET untuk /
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Menangani rute GET untuk /hello/:name
    app.Get("/hello/:name", func(c *fiber.Ctx) error {
        name := c.Params("name")
        return c.SendString("Hello, " + name + "!")
    })

    // Menangani rute POST untuk /data
    app.Post("/data", func(c *fiber.Ctx) error {
        type Data struct {
            Name string `json:"name"`
            Age  int    `json:"age"`
        }

        var data Data

        // Mengurai JSON dari request body
        if err := c.BodyParser(&data); err != nil {
            return c.Status(400).SendString("Bad Request")
        }

        return c.JSON(data)
    })

    // Menjalankan server di port 3000
    err := app.Listen(":3000")
    if err != nil {
        panic(err)
    }
}
