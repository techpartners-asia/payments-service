package adapters

import (
	"fmt"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	sharedDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/shared"
	hipaygo "github.com/techpartners-asia/hipay-go/hipay"
)

type HipayAdapter struct {
	client hipaygo.Hipay
}

func NewHipayAdapter(input sharedDTO.HipayAdapterDTO) *HipayAdapter {
	return &HipayAdapter{client: hipaygo.New(input.Endpoint, input.Token, input.EntityID)}
}

func (a *HipayAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("hipay adapter not configured")
	}

	res, err := a.client.Checkout(payment.Amount)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: res.CheckoutID,
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *HipayAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("hipay adapter not configured")
	}

	res, err := a.client.PaymentGet(payment.RefInvoiceID)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: res.ResultCode == "COMPLETED",
	}, nil
}
