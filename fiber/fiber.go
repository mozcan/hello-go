package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	//middleware
	/*
		app.Use(func(ctx *fiber.Ctx) error {
			fmt.Printf("Hello client. You're calling my -> %s%s and method %s\n", ctx.BaseURL(), ctx.Request().RequestURI(), ctx.Request().Header.Method())
			return ctx.Next()
		})

	*/

	//middleware for spesific endponits
	app.Use("/user", func(ctx *fiber.Ctx) error {
		correlationId := ctx.Get("X-CorrelationId")

		if correlationId == "" {
			return ctx.Status(http.StatusBadRequest).JSON("You have to send correlationId")
		}

		_, err := uuid.Parse(correlationId)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON("CorrelationId must be guid value")
		}

		return ctx.Next()
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		fmt.Println("Hello First Get Endpoint")
		ctx.Status(http.StatusOK)
		return ctx.JSON("Hello First get endpoint")
	})

	app.Get("/user/:userId", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("userId")
		fmt.Sprintf("User Id -> %s\n", userId)
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
