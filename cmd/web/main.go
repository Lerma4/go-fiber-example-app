package main

import (
	components "go-fiber/ui/html"
	"log/slog"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type application struct {
	logger *slog.Logger
}

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	server := fiber.New()

	app.logger.Info("Starting server")

	// Use the Static method to serve static files such as images, CSS, and JavaScript.
	server.Static("/", "./ui/static")

	server.Get("/login", func(c *fiber.Ctx) error {
		return Render(c, components.Login())
	})

	// Definisci il gestore di route per l'endpoint "/submit"
	server.Post("/submit", func(c *fiber.Ctx) error {
		// Ottieni i dati dal modulo inviato
		email := c.FormValue("email")

		// Esempio di elaborazione dei dati (in questo caso, restituiamo un messaggio di conferma)
		message := "Grazie" + "! Il tuo indirizzo email " + email + " è stato registrato correttamente."

		// Restituisci la risposta
		return c.SendString(message)
	})

	server.Get("/test/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		if name == "" {
			name = "World"
		}
		return Render(c, components.Home(name))
	})
	server.Use(NotFoundMiddleware)

	log.Fatal(server.Listen(":3000"))
}

func NotFoundMiddleware(c *fiber.Ctx) error {
	return Render(c, components.NotFound(), templ.WithStatus(http.StatusNotFound))
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
