package services

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/OzkrOssa/redplanet-bridge/internal/models"
	"github.com/OzkrOssa/redplanet-bridge/internal/repository"
)

type PaymentezService struct {
	repo *repository.PaymentezRepository
}

func NewPaymentezService(repo *repository.PaymentezRepository) *PaymentezService {
	return &PaymentezService{repo}
}

func (s *PaymentezService) hastData(appKey, appCode string) string {
	unixTimeStamp := time.Now().Unix()
	TokenString := appKey + fmt.Sprint(unixTimeStamp)

	hash := sha256.New()
	hash.Write([]byte(TokenString))
	tokenHash := hex.EncodeToString(hash.Sum(nil))

	token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s;%d;%s", appCode, unixTimeStamp, tokenHash)))
	return token
}

func (s *PaymentezService) GenerateToken(payMethod string) string {
	switch payMethod {
	case "pse":
		return s.hastData(os.Getenv("SERVER_PSE_APPKEY"), os.Getenv("SERVER_PSE_APPCODE"))
	case "card":
		return s.hastData(os.Getenv("SERVER_CARD_APPKEY"), os.Getenv("SERVER_CARD_APPCODE"))
	default:
		return "invalid pay method must be (pse/card)"
	}
}

func (s *PaymentezService) PsePaymentWithSplits(data *models.PaymentRequetsPayload) (interface{}, error) {
	payload := map[string]interface{}{
		"carrier": map[string]interface{}{
			"id": "PSE",
			"extra_params": map[string]interface{}{
				"bank_code":    "1022",
				"response_url": "https://thanks.red-planet.com.co/",
				"user": map[string]interface{}{
					"name":            data.User.Name,
					"fiscal_number":   data.User.FiscalNumber,
					"type":            data.User.Type,
					"type_fis_number": data.User.TypeFisNumber,
					"ip_address":      data.IPAddress,
				},
				"split": map[string]interface{}{
					"transactions": []map[string]interface{}{
						{
							"application_code": os.Getenv("SPLIT_H"),
							"amount":           data.Order.Amount * 0.5,
							"vat":              data.Order.Vat,
						},
						{
							"application_code": os.Getenv("SPLIT_H2"),
							"amount":           data.Order.Amount * 0.2,
							"vat":              data.Order.Vat,
						},
						{
							"application_code": os.Getenv("SPLIT_H3"),
							"amount":           data.Order.Amount * 0.3,
							"vat":              data.Order.Vat,
						},
					},
				},
			},
		},
		"user": map[string]interface{}{
			"id":           data.User.ID,
			"email":        data.User.Email,
			"phone_number": data.User.PhoneNumber,
		},
		"order": map[string]interface{}{
			"dev_reference": data.Order.DevReference,
			"amount":        data.Order.Amount,
			"vat":           data.Order.Vat,
			"description":   data.Order.Description,
		},
	}

	encodedPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, "https://noccapi-stg.paymentez.com/order/", bytes.NewBuffer(encodedPayload))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Auth-Token", s.GenerateToken("pse"))

	client := &http.Client{
		// Transport: &http.Transport{
		// 	TLSClientConfig: &tls.Config{
		// 		InsecureSkipVerify: true,
		// 	},
		// },
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		byteResponse, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf(string(byteResponse))
	}

	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var bindResponse models.PsePaymentResponse

	err = json.Unmarshal(byteResponse, &bindResponse)
	if err != nil {
		return nil, err
	}

	return bindResponse, nil
}

func (s *PaymentezService) ProcessEventWebHook(event *models.WebhookEvent) error {
	err := ValidateStoken(event)
	if err != nil {
		return err
	}

	s.repo.StoreWebhook(event)

	return nil
}
