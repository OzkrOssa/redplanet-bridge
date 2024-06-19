package handlers

import (
	"net/http"

	"github.com/OzkrOssa/redplanet-bridge/internal/models"
	"github.com/OzkrOssa/redplanet-bridge/internal/services"
	"github.com/labstack/echo/v5"
)

type PaymentezHandler struct {
	ps services.PaymentezService
}

func NewPaymentezHandler() *PaymentezHandler {
	return &PaymentezHandler{}
}

func (ph *PaymentezHandler) GenerateToken(c echo.Context) error {
	payMethod := c.PathParam("payMethod")
	if payMethod != "card" && payMethod != "pse" {
		return c.JSON(400, map[string]interface{}{
			"status":   http.StatusBadRequest,
			"message":  "invalid param, must be card or pse",
			"response": nil,
		})
	}

	token := ph.ps.GenerateToken(payMethod)

	return c.JSON(200, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "token generate successfully",
		"response": map[string]interface{}{
			"token": token,
		},
	})
}

func (ph *PaymentezHandler) PsePaymentWithSplit(c echo.Context) error {

	p := new(models.PaymentRequetsPayload)

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   http.StatusInternalServerError,
			"message":  err.Error(),
			"response": nil,
		})
	}

	paymentResponse, err := ph.ps.PsePaymentWithSplits(p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":   http.StatusBadRequest,
			"message":  err.Error(),
			"response": nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   http.StatusOK,
		"message":  "pse payment successfully",
		"response": paymentResponse,
	})
}
