package routes

import (
	"os"

	"github.com/OzkrOssa/redplanet-bridge/internal/handlers"
	"github.com/OzkrOssa/redplanet-bridge/internal/repository"
	"github.com/OzkrOssa/redplanet-bridge/internal/services"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type PaymentezRouter struct{}

func (t *PaymentezRouter) V1PaymentezRoutes(e *core.ServeEvent, app *pocketbase.PocketBase) {
	group := e.Router.Group("api/v1/paymentez")

	group.Use(middleware.KeyAuth(func(c echo.Context, key string, source middleware.ExtractorSource) (bool, error) {
		return key == os.Getenv("BEARER_TOKEN"), nil
	}))

	paymentezHandler := handlers.NewPaymentezHandler(app, services.NewPaymentezService(repository.NewPaymentezRepository(app)))

	group.GET("/token/:payMethod", echo.HandlerFunc(func(c echo.Context) error {
		return paymentezHandler.GenerateToken(c)
	}))

	group.POST("/pse/split", echo.HandlerFunc(func(c echo.Context) error {
		return paymentezHandler.PsePaymentWithSplit(c)
	}))

	group.POST("/webhook", echo.HandlerFunc(func(c echo.Context) error {
		return paymentezHandler.ProcessEventWebHook(c)
	}))
}
