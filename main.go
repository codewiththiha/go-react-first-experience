package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TODO struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello World")
	app := fiber.New()

	todos := []TODO{}

	// Listen the output at the home route of host "/"
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// * fiber.Ctx is used as a communicator between user and responses from the api
	app.Post("/api/todos/", func(c *fiber.Ctx) error {
		todo := &TODO{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "body required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)

	})

	// Update a TODO
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		// Since this create copies
		// for _, todo := range todos {
		// 	if fmt.Sprint(todo.ID) == id {
		// 		todo.Completed = true
		// 		return c.Status(200).JSON(todo)
		// 	}
		// }

		// This actually modify the real array slices since it uses index

		// The first snippet would only work if todos was a slice of pointers (e.g., []*Todo). In that case,
		// todo would be a copy of the pointer, but both the copy and the original would point to the same memory address,
		// allowing the update to work.

		for i := range todos {
			if fmt.Sprint(todos[i].ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not Found"})

	})

	// delete a todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	})

	log.Fatal(app.Listen(":4000"))
}
