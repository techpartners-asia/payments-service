package adapters

import (
	"fmt"

	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	paymentServiceResponseDTO "github.com/techpartners-asia/payments-service/infrastructure/payment/dto/response"
	sharedDTO "github.com/techpartners-asia/payments-service/infrastructure/shared"

	golomt "github.com/techpartners-asia/golomt-api-go/ecommerce"
)

// GolomtAdapter implements PaymentProvider for Golomt ecommerce.
type GolomtAdapter struct {
	client golomt.GolomtEcommerce
	input  sharedDTO.GolomtAdapterDTO
}

func NewGolomtAdapter(input sharedDTO.GolomtAdapterDTO) *GolomtAdapter {
	return &GolomtAdapter{client: golomt.New(input.BaseURL, input.Secret, input.BearerToken)}
}

func (a *GolomtAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("golomt adapter not configured")
	}

	returnType := golomt.GET
	if a.input.ReturnType != "" {
		switch a.input.ReturnType {
		case "GET", "get":
			returnType = golomt.GET
		case "POST", "post":
			returnType = golomt.POST
		case "MOBILE", "mobile":
			returnType = golomt.MOBILE
		default:
			return nil, fmt.Errorf("invalid golomt return type: %s", a.input.ReturnType)
		}
	}

	req := golomt.CreateInvoiceInput{
		ReturnType:    returnType,
		Amount:        payment.Amount,
		TransactionID: payment.UID,
		Callback:      a.input.CallbackURL,
	}

	res, err := a.client.CreateInvoice(req)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: res.Invoice,
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *GolomtAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("golomt adapter not configured")
	}

	res, err := a.client.Inquiry(payment.UID)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: res.Status == "000",
	}, nil
}
