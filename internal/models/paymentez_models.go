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
