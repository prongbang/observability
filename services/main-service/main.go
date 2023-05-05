package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prongbang/main-service/proto/auth"
	"github.com/prongbang/main-service/proto/coin"
	"github.com/prongbang/main-service/proto/user"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	userAddress = "user_service:50051"
	authAddress = "auth_service:50052"
	coinAddress = "coin_service:50053"
)

func main() {

	// Set up a connection to the server
	authConn, authErr := grpc.Dial(authAddress, grpc.WithInsecure())
	userConn, userErr := grpc.Dial(userAddress, grpc.WithInsecure())
	coinConn, coinErr := grpc.Dial(coinAddress, grpc.WithInsecure())
	if authErr != nil {
		log.Fatalf("did not connect: %v", authErr)
	}
	if userErr != nil {
		log.Fatalf("did not connect: %v", userErr)
	}
	if coinErr != nil {
		log.Fatalf("did not connect: %v", coinErr)
	}
	defer func(authConn *grpc.ClientConn) { _ = authConn.Close() }(authConn)
	defer func(userConn *grpc.ClientConn) { _ = userConn.Close() }(userConn)
	defer func(coinConn *grpc.ClientConn) { _ = coinConn.Close() }(coinConn)
	authClient := auth.NewAuthClient(authConn)
	userClient := user.NewUserClient(userConn)
	coinClient := coin.NewCoinClient(coinConn)

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
				return c.JSON(fiber.Map{"message": err.Error()})
			}
			return c.JSON(resp)
		})

		v1.Get("/user", func(c *fiber.Ctx) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			resp, err := userClient.GetUser(ctx, &user.UserRequest{Username: "em"})
			if err != nil {
				return c.JSON(fiber.Map{"message": err.Error()})
			}
			return c.JSON(resp)
		})

		v1.Get("/coin", func(c *fiber.Ctx) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			resp, err := coinClient.GetCoin(ctx, &coin.CoinRequest{Username: "em"})
			if err != nil {
				return c.JSON(fiber.Map{"message": err.Error()})
			}
			return c.JSON(resp)
		})
	}

	log.Fatal(app.Listen(":8000"))
}
