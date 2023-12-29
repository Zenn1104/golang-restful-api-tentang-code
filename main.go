package main

import (
	"restful-api-golang-tentang-code/controllers/bookcontroller"
	"restful-api-golang-tentang-code/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/book")

	book.Get("/", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("/", bookcontroller.Create)
	book.Put("/:id", bookcontroller.Update)
	book.Delete(":id", bookcontroller.Delete)

	app.Listen(":8000")
}
