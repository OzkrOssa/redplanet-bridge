package services

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/OzkrOssa/redplanet-bridge/internal/models"
)

func ValidateStoken(event *models.WebhookEvent) error {
	var appCode, appKey string
	var temp struct {
		Split *models.WebHookSplit
		User  *models.WebHookUser
		Card  *models.WebHookCard
	}

	whByte, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(whByte, &temp); err != nil {
		return err
	}

	switch {
	case temp.Split != nil:
		appCode = os.Getenv("SERVER_PSE_APPCODE")
		appKey = os.Getenv("SERVER_PSE_APPKEY")
	case temp.Card != nil:
		appCode = os.Getenv("SERVER_CARD_APPCODE")
		appKey = os.Getenv("SERVER_CARD_APPKEY")
	default:
		return fmt.Errorf("missing required fields")
	}

	hashPayload := fmt.Sprintf("%s_%s_%s_%s", event.Transaction.ID, appCode, event.User.ID, appKey)
	hash := md5.New()
	hash.Write([]byte(hashPayload))
	tokenHash := hex.EncodeToString(hash.Sum(nil))

	if tokenHash != event.Transaction.Stoken {
		return fmt.Errorf("the provided stoken does not match the expected stoken")
	}

	return nil
}
