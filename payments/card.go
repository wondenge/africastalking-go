package atgo

import (
	"context"
)

type (
	Card interface {

		// Collect money into your Payment Wallet by initiating transactions that deduct
		// money from a customers Debit or Credit Card.
		CardCheckout(ctx context.Context, p *CardCheckoutPayload) (res *CardCheckoutResponse, err error)

		// Allows your application to validate card checkout charge requests.
		CardCheckoutValidate(ctx context.Context, p *CardCheckoutValidatePayload) (res *CardCheckoutValidateResponse, err error)
	}
)

// CardCheckoutPayload is the payload type of the africastalking service
// CardCheckout method.
type CardCheckoutPayload struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// Africa’s Talking Payment Product initiating transaction.
	ProductName string      `form:"productName" json:"productName" xml:"productName"`
	PaymentCard PaymentCard `form:"paymentCard,omitempty" json:"paymentCard,omitempty" xml:"paymentCard,omitempty"`
	// Token generated by AfricasTalking
	CheckoutToken string `form:"checkoutToken,omitempty" json:"checkoutToken,omitempty" xml:"checkoutToken,omitempty"`
	// 3-digit ISO format currency code.
	CurrencyCode string `form:"currencyCode" json:"currencyCode" xml:"currencyCode"`
	// Amount client is expected to confirm.
	Amount float64 `form:"amount" json:"amount" xml:"amount"`
	// Short description of the transaction.
	Narration string `form:"narration" json:"narration" xml:"narration"`
	// Metadata associated with the request.
	Metadata map[string]string `form:"metadata,omitempty" json:"metadata,omitempty" xml:"metadata,omitempty"`
}

// CardCheckoutResponse is the result type of the africastalking service
// CardCheckout method.
type CardCheckoutResponse struct {
	// The status of the request.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A detailed description of the request status.
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Unique ID generated for successful requests.
	TransactionID string `form:"transactionId,omitempty" json:"transactionId,omitempty" xml:"transactionId,omitempty"`
}

// Details of the payment card to be charged.
type PaymentCard struct {
	// Payment card number.
	Number string `form:"number" json:"number" xml:"number"`
	// 3 or 4 digit payment card verification value.
	CvvNumber int `form:"cvvNumber" json:"cvvNumber" xml:"cvvNumber"`
	// Expiration month on the payment card.
	ExpiryMonth int `form:"expiryMonth" json:"expiryMonth" xml:"expiryMonth"`
	// Expiration year on the payment card.
	ExpiryYear int `form:"expiryYear" json:"expiryYear" xml:"expiryYear"`
	// 2-Digit country code where the payment card was issued.
	CountryCode string `form:"countryCode" json:"countryCode" xml:"countryCode"`
	// The payment cards ATM PIN.
	AuthToken string `form:"authToken" json:"authToken" xml:"authToken"`
}

// CardCheckoutValidatePayload is the payload type of the africastalking
// service CardCheckoutValidate method.
type CardCheckoutValidatePayload struct {
	// Africa’s Talking application username.
	Username string `form:"username" json:"username" xml:"username"`
	// ID of the transaction application wants to validate.
	TransactionID string `form:"transactionId" json:"transactionId" xml:"transactionId"`
	// One Time Password card provider sent to the client.
	Otp string `form:"otp" json:"otp" xml:"otp"`
}

// CardCheckoutValidateResponse is the result type of the africastalking
// service CardCheckoutValidate method.
type CardCheckoutValidateResponse struct {
	// Corresponds to the final status of this request.
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// A detailed description of the request status.
	Description string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Token application can use to initiate subsequent charges.
	CheckoutToken string `form:"checkoutToken,omitempty" json:"checkoutToken,omitempty" xml:"checkoutToken,omitempty"`
}
