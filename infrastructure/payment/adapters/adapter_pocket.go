package adapters

import (
	"fmt"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	sharedDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/shared"

	pocket "github.com/techpartners-asia/pocket-go"
)

// PocketAdapter implements PaymentProvider for Pocket.
type PocketAdapter struct {
	client pocket.Pocket
}

func NewPocketAdapter(input sharedDTO.PocketAdapterDTO) *PocketAdapter {
	return &PocketAdapter{client: pocket.New(input.Merchant, input.ClientID, input.ClientSecret, input.Environment, input.TerminalIDRaw)}
}

func (a *PocketAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("pocket adapter not configured")
	}

	req := pocket.PocketCreateInvoiceInput{
		Amount:      payment.Amount,
		OrderNumber: payment.UID,
		InvoiceType: "ZERO",
		Channel:     "merchant",
		Info:        payment.Note,
	}

	res, err := a.client.CreateInvoice(req)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: fmt.Sprintf("%d", res.ID),
		BankQRCode:    res.Qr,
		Deeplinks: []paymentServiceResponseDTO.Deeplink{{
			Name:        "Pocket",
			Description: "Pocket",
			Link:        res.DeepLink,
		}},
		IsPaid: false,
		Raw:    res,
	}, nil
}

func (a *PocketAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("pocket adapter not configured")
	}

	res, err := a.client.GetInvoiceByOrderNumber(payment.UID)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: res.State == "paid",
	}, nil
}
