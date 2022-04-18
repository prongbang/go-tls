package main

import (
	"crypto/tls"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"protocol": c.Protocol(),
		})
	})

	// Create tls certificate
	cer, err := tls.LoadX509KeyPair("mkcert/tls.go.com+4.pem", "mkcert/tls.go.com+4-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	// Create custom listener
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		panic(err)
	}

	log.Fatal(app.Listener(ln))
}
