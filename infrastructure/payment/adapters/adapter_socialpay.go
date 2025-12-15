package adapters

import (
	"fmt"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	sharedDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/shared"

	"github.com/techpartners-asia/golomt-api-go/socialpay"
)

// SocialPayAdapter implements PaymentProvider for SocialPay.
type SocialPayAdapter struct {
	client socialpay.SocialPay
}

func NewSocialPayAdapter(input sharedDTO.SocialPayAdapterDTO) *SocialPayAdapter {
	return &SocialPayAdapter{client: socialpay.New(input.Terminal, input.Secret, input.Endpoint)}
}

func (a *SocialPayAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("socialpay adapter not configured")
	}

	res, err := a.client.CreateInvoiceQR(socialpay.InvoiceInput{
		Amount:  payment.Amount,
		Invoice: payment.UID,
	})
	if err != nil {
		return nil, err
	}

	// The library response does not include a QR text field; return raw and invoice id.
	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: payment.UID,
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *SocialPayAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("socialpay adapter not configured")
	}

	res, err := a.client.CheckInvoice(socialpay.InvoiceInput{
		Invoice: payment.UID,
		Amount:  payment.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: res.ResponseCode == "00",
	}, nil
}
