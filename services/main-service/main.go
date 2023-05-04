package main

import (
	"context"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prongbang/main-service/proto/auth"
	"github.com/prongbang/main-service/proto/user"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

const (
	authAddress = "localhost:50052"
	userAddress = "localhost:50051"
)

func main() {

	// Set up a connection to the server
	authConn, authErr := grpc.Dial(authAddress, grpc.WithInsecure())
	userConn, userErr := grpc.Dial(userAddress, grpc.WithInsecure())
	if authErr != nil {
		log.Fatalf("did not connect: %v", authErr)
	}
	if userErr != nil {
		log.Fatalf("did not connect: %v", userErr)
	}
	defer func(authConn *grpc.ClientConn) { _ = authConn.Close() }(authConn)
	defer func(userConn *grpc.ClientConn) { _ = userConn.Close() }(userConn)
	authClient := auth.NewAuthClient(authConn)
	userClient := user.NewUserClient(userConn)

	// New creates a new Fiber named instance.
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Routers
	v1 := app.Group("/v1")
	{
		v1.Get("/login", func(c *fiber.Ctx) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			resp, err := authClient.Login(ctx, &auth.LoginRequest{Username: "em", Password: "1234"})
			if err != nil {
				return fiber.ErrBadRequest
			}
			return c.JSON(resp)
		})

		v1.Get("/user", func(c *fiber.Ctx) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			resp, err := userClient.GetUser(ctx, &user.UserRequest{Username: "em"})
			if err != nil {
				return fiber.ErrBadRequest
			}
			return c.JSON(resp)
		})

		v1.Get("/coin", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})
	}

	log.Fatal(app.Listen(":3000"))
}
