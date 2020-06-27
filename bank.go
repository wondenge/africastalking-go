package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

const (

	// Base URLs:
	BankCheckoutLiveURL    = "https://payments.africastalking.com/bank/checkout/charge"
	BankCheckoutSandboxURL = "https://payments.sandbox.africastalking.com/bank/checkout/charge"

	BankCheckoutValidateLiveURL    = "https://payments.africastalking.com/bank/checkout/validate"
	BankCheckoutValidateSandboxURL = "https://payments.sandbox.africastalking.com/bank/checkout/validate"

	BankTransferLiveURL    = "https://payments.africastalking.com/bank/transfer"
	BankTransferSandboxURL = "https://payments.sandbox.africastalking.com/bank/transfer"

	// Supported Banks:
	FCMBNigeria       = 234001
	ZenithNigeria     = 234002
	AccessNigeria     = 234003
	GTBankNigeria     = 234004
	EcobankNigeria    = 234005
	DiamondNigeria    = 234006
	ProvidusNigeria   = 234007
	UnityNigeria      = 234008
	StanbicNigeria    = 234009
	SterlingNigeria   = 234010
	ParkwayNigeria    = 234011
	AfribankNigeria   = 234012
	EnterpriseNigeria = 234013
	FidelityNigeria   = 234014
	HeritageNigeria   = 234015
	KeystoneNigeria   = 234016
	SkyeNigeria       = 234017
	StanchartNigeria  = 234018
	UnionNigeria      = 234019
	UBANigeria        = 234020
	WemaNigeria       = 234021
	FirstNigeria      = 234022
)

// Collect money into your payment wallet.
func (c *Client) bankCheckout(ctx context.Context, p *africastalking.BankCheckoutPayload) (res *africastalking.BankCheckoutResponse, err error) {

	req, err := c.NewRequest("POST", "https://payments.sandbox.africastalking.com/bank/checkout/charge", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")
	res = &africastalking.BankCheckoutResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Validate a bank checkout charge request
func (c *Client) bankCheckoutValidate(ctx context.Context, p *africastalking.BankCheckoutValidatePayload) (res *africastalking.BankCheckoutValidateResponse, err error) {

	req, err := c.NewRequest("POST", "https://payments.sandbox.africastalking.com/bank/checkout/validate", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.BankCheckoutValidateResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Initiate a bank transfer request.
func (c *Client) bankTransfer(ctx context.Context, p *africastalking.BankTransferPayload) (res *africastalking.BankTransferResponse, err error) {

	req, err := c.NewRequest("POST", "https://payments.sandbox.africastalking.com/bank/transfer", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.BankTransferResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}