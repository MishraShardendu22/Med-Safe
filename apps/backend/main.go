package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func init(){
	fmt.Println("Initializing Med-Safe Backend...")


}

func main() {
	fmt.Println("Hello, Med-Safe!")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Med-Safe Backend!")
	})

	app.Listen(":5000")
	fmt.Println("Med-Safe Backend is ready!")
}