package repository

import (
	"encoding/json"
	"log/slog"

	"github.com/OzkrOssa/redplanet-bridge/internal/models"
	"github.com/OzkrOssa/redplanet-bridge/internal/utils"
	"github.com/pocketbase/pocketbase"
	pm "github.com/pocketbase/pocketbase/models"
)

type PaymentezRepository struct {
	App *pocketbase.PocketBase
}

func NewPaymentezRepository(App *pocketbase.PocketBase) *PaymentezRepository {
	return &PaymentezRepository{App}
}

func (pr *PaymentezRepository) StoreWebhook(event *models.WebhookEvent) {

	var temp struct {
		Split *models.WebHookSplit
		User  *models.WebHookUser
		Card  *models.WebHookCard
	}

	eventByte, err := json.Marshal(event)
	if err != nil {
		slog.Error(
			"error",
			"error to marshal evento", err.Error(),
		)
		return
	}

	if err := json.Unmarshal(eventByte, &temp); err != nil {
		slog.Error(
			"error",
			"error Unmarshal eventByte to temp struct", err.Error(),
		)
		return
	}
	tID, _ := pr.App.Dao().FindFirstRecordByData("webhooks", "paymentez_id", event.Transaction.ID)

	if tID != nil && tID.Id != "" {
		backupCollection, err := pr.App.Dao().FindCollectionByNameOrId("duplicate_webhooks")
		if err != nil {
			slog.Error(
				"error",
				"collection duplicate_webhooks not found", err.Error(),
			)
			return
		}

		backupRecord := pm.NewRecord(backupCollection)
		backupRecord.Set("paymentez_id", event.Transaction.ID)
		backupRecord.Set("data", eventByte)

		if err := pr.App.Dao().SaveRecord(backupRecord); err != nil {
			slog.Error(
				"error",
				"error to save record in collection (duplicate_webhook)", err.Error(),
			)
			return
		}
		return
	}

	// Si el w_id no existe, guarda el nuevo webhook
	collection, err := pr.App.Dao().FindCollectionByNameOrId("webhooks")
	if err != nil {
		slog.Error(
			"error",
			"collection webhooks not found", err.Error(),
		)
		return
	}

	record := pm.NewRecord(collection)

	if temp.Card != nil {
		record.Set("pay_method", "card")
	}

	if temp.Split != nil {
		record.Set("pay_method", "pse")
	}

	record.Set("paymentez_id", event.Transaction.ID)
	record.Set("status", utils.STATUS[event.Transaction.Status])
	record.Set("status_detail", utils.STATUS_DETALS[event.Transaction.StatusDetail])
	record.Set("order_description", event.Transaction.OrderDescription)
	record.Set("authorization_code", event.Transaction.AuthorizationCode)
	record.Set("date", event.Transaction.Date)
	record.Set("dev_reference", event.Transaction.DevReference)
	record.Set("amount", event.Transaction.Amount)
	record.Set("paid_date", event.Transaction.PaidDate)
	record.Set("stoken", event.Transaction.Stoken)
	record.Set("application_code", event.Transaction.ApplicationCode)
	record.Set("user_id", event.User.ID)
	record.Set("user_email", event.User.Email)

	if err := pr.App.Dao().SaveRecord(record); err != nil {
		slog.Error(
			"error",
			"error to save record on collection webhook", err.Error(),
		)
		return
	}
}
