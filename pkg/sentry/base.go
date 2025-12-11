package sentryPkg

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/fiber/v2"
)

func InitializeSentry() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://e5cb4aef9b44dcdda6f09933c4d0ff61@o4510458762100736.ingest.de.sentry.io/4510458776059989",
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
}

func NewSentry() fiber.Handler {
	return sentryfiber.New(sentryfiber.Options{
		Repanic:         true,
		WaitForDelivery: true,
	})
}
