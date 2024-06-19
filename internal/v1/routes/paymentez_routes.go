package routes

import (
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase/core"
	"guithub.com/OzkrOssa/redplanet-bridge/internal/handlers"
)

type PaymentezRouter struct{}

func (t *PaymentezRouter) V1PaymentezRoutes(e *core.ServeEvent) {
	group := e.Router.Group("api/v1/paymentez")

	group.Use(middleware.KeyAuth(func(c echo.Context, key string, source middleware.ExtractorSource) (bool, error) {
		return key == os.Getenv("BEARER_TOKEN"), nil
	}))

	paymentezHandler := handlers.NewPaymentezHandler()

	group.GET("/token/:payMethod", echo.HandlerFunc(func(c echo.Context) error {
		return paymentezHandler.GenerateToken(c)
	}))

	group.POST("/pse/split", echo.HandlerFunc(func(c echo.Context) error {
		return paymentezHandler.PsePaymentWithSplit(c)
	}))
}
