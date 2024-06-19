package models

type CreatePsePaymentRequestToPaymentez struct {
	Carrier Carrier     `json:"carrier"`
	User    UserDetails `json:"user"`
	Order   Order       `json:"order"`
}

type Carrier struct {
	ID          string      `json:"id"`
	ExtraParams ExtraParams `json:"extra_params"`
}

type ExtraParams struct {
	BankCode    string  `json:"bank_code"`
	ResponseURL string  `json:"response_url"`
	User        PseUser `json:"user"`
}

type PseUser struct {
	Name          string `json:"name"`
	FiscalNumber  int64  `json:"fiscal_number"`
	Type          string `json:"type"`
	TypeFisNumber string `json:"type_fis_number"`
	IPAddress     string `json:"ip_address"`
}

type UserDetails struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type Order struct {
	DevReference string  `json:"dev_reference"`
	Amount       float64 `json:"amount"`
	Vat          float64 `json:"vat"`
	Description  string  `json:"description"`
}

//------------------- RESPONSE--------------------------//

type PsePaymentResponse struct {
	Application Application         `json:"application"`
	Commerce    Commerce            `json:"commerce"`
	User        PaymentResponseUser `json:"user"`
	Transaction Transaction         `json:"transaction"`
}

type Application struct {
	Code string `json:"code"`
}

type Commerce struct {
	MerchantID string `json:"merchant_id"`
}

type PaymentResponseUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
}

type Transaction struct {
	Currency        string  `json:"currency"`
	Country         string  `json:"country"`
	DevReference    string  `json:"dev_reference"`
	Amount          float64 `json:"amount"`
	PaidDate        *string `json:"paid_date"` // Usamos un puntero para poder manejar valores nulos
	Description     string  `json:"description"`
	Status          string  `json:"status"`
	ID              string  `json:"id"`
	BankURL         string  `json:"bank_url"`
	StatusBank      string  `json:"status_bank"`
	TrazabilityCode string  `json:"trazability_code"`
	TicketID        int     `json:"ticket_id"`
	PSECycle        string  `json:"pse_cycle"`
	Bank            Bank    `json:"bank,omitempty"`
}

type Bank struct {
	Name string `json:"name,omitempty"`
	Code int    `json:"code,omitempty"`
}

type PaymentRequetsPayload struct {
	User struct {
		ID            string `json:"id,omitempty"`
		Name          string `json:"name,omitempty"`
		FiscalNumber  int    `json:"fiscal_number,omitempty"`
		Type          string `json:"type,omitempty"`
		TypeFisNumber string `json:"type_fis_number,omitempty"`
		Email         string `json:"email,omitempty"`
		PhoneNumber   string `json:"phone_number,omitempty"`
	} `json:"user,omitempty"`
	Order     Order  `json:"order"`
	IPAddress string `json:"ip_address"`
}

// -----------WebHook---------------------
type WebhookEvent struct {
	Transaction WebHookTransaction `json:"transaction,omitempty"`
	Split       *WebHookSplit      `json:"split,omitempty"`
	User        *WebHookUser       `json:"user,omitempty"`
	Card        *WebHookCard       `json:"card,omitempty"`
}

type WebHookTransaction struct {
	Status            string `json:"status,omitempty"`
	OrderDescription  string `json:"order_description,omitempty"`
	AuthorizationCode string `json:"authorization_code,omitempty"`
	StatusDetail      string `json:"status_detail,omitempty"`
	Date              string `json:"date,omitempty"`
	Message           string `json:"message,omitempty"`
	ID                string `json:"id,omitempty"`
	DevReference      string `json:"dev_reference,omitempty"`
	CarrierCode       string `json:"carrier_code,omitempty"`
	Amount            string `json:"amount,omitempty"`
	PaidDate          string `json:"paid_date,omitempty"`
	Installments      string `json:"installments,omitempty"`
	LtpID             string `json:"ltp_id,omitempty"`
	Stoken            string `json:"stoken,omitempty"`
	ApplicationCode   string `json:"application_code,omitempty"`
	TerminalCode      string `json:"terminal_code,omitempty"`
}

type WebHookUser struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type WebHookCard struct {
	Bin          string `json:"bin,omitempty"`
	HolderName   string `json:"holder_name,omitempty"`
	Type         string `json:"type,omitempty"`
	Number       string `json:"number,omitempty"`
	Origin       string `json:"origin,omitempty"`
	FiscalNumber string `json:"fiscal_number,omitempty"`
}

type WebHookSplit struct {
	Transaction []WebHookSplitTransactions `json:"transaction,omitempty"`
}

type WebHookSplitTransactions struct {
	Installments      string `json:"installments,omitempty"`
	Amount            string `json:"amount,omitempty"`
	ID                string `json:"id,omitempty"`
	ApplicationCode   string `json:"application_code,omitempty"`
	AuthorizationCode string `json:"authorization_code,omitempty"`
}
