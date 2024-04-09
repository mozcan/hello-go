package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserCreateRequwst struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int32  `json:"age"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		fmt.Println("Hello First Get Endpoint")
		ctx.Status(http.StatusOK)
		return ctx.JSON("Hello First get endpoint")
	})

	app.Get("/user/:userId", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId")
		fmt.Sprintf("User Id -> %s", userId)
		return ctx.SendString(fmt.Sprintf("User Id -> %s", userId))
	})

	app.Post("/user", func(ctx *fiber.Ctx) error {
		fmt.Println("This is my first Post request")
		var request UserCreateRequwst

		err := ctx.BodyParser(&request)
		if err != nil {
			fmt.Sprintf("There was an error while binding rquest -> Error: %v", err.Error())
			return err
		}

		return ctx.Status(http.StatusOK).JSON(fmt.Sprintf("%s created successfuly", request.FirstName))
	})

	app.Listen(":3000")
}
