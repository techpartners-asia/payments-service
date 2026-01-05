package adapters

import (
	"fmt"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	sharedDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/shared"
	bpaygo "github.com/techpartners-asia/bpay-go"
)

type BpayAdapter struct {
	client bpaygo.Bpay
}

func NewBpayAdapter(input sharedDTO.BpayAdapterDTO) *BpayAdapter {
	return &BpayAdapter{client: bpaygo.New(input.Endpoint, input.Username, input.Password)}
}

func (a *BpayAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("bpay adapter not configured")
	}

	res, err := a.client.InvoiceCreate(bpaygo.BpayInvoiceCreateRequest{
		BillIDs: []int64{},
	}, 1)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: fmt.Sprintf("%d", res.ID),
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *BpayAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("bpay adapter not configured")
	}

	res, err := a.client.BillCheck(payment.UID)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: res.Status == "SUCCESS",
	}, nil
}
