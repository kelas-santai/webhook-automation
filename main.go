package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Get PORT and SECRET_KEY from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Println("WARNING: SECRET_KEY not set")
	}

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, world!"})
	})

	app.Post("/webhook-backend", func(c *fiber.Ctx) error {
		event := c.Get("X-GitHub-Event")
		var payload map[string]interface{}
		if err := c.BodyParser(&payload); err != nil {
			return fiber.ErrBadRequest
		}

		go handleWebhook(event, payload, "backend")

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": "Accepted"})
	})

	app.Post("/webhook-frontend", func(c *fiber.Ctx) error {
		event := c.Get("X-GitHub-Event")
		var payload map[string]interface{}
		if err := c.BodyParser(&payload); err != nil {
			return fiber.ErrBadRequest
		}

		go handleWebhook(event, payload, "frontend")

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": "Accepted"})
	})

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func handleWebhook(event string, data map[string]interface{}, target string) {
	switch event {
	case "ping":
		fmt.Println("GitHub sent the ping event")
	case "issues":
		action := data["action"]
		issue := data["issue"].(map[string]interface{})
		switch action {
		case "opened":
			fmt.Println("An issue was opened with this title:", issue["title"])
		case "closed":
			user := issue["user"].(map[string]interface{})
			fmt.Println("An issue was closed by", user["login"])
		default:
			fmt.Println("Unhandled action for the issue event:", action)
		}
	case "push":
		ref := data["ref"]
		if ref == "refs/heads/main" {
			if target == "backend" {
				fmt.Println("Triggering backend deployment...")
				exec.Command("/home/kelas-santai/webhooks/webhooks/deploy.sh").Start()
			} else if target == "frontend" {
				fmt.Println("Triggering frontend deployment...")
				exec.Command("/home/kelas-santai/webhooks/webhooks/deploy-fe.sh").Start()
			}
		}
	default:
		fmt.Println("Unhandled event:", event)
	}
}
