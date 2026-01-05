package adapters

import (
	"fmt"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	adapterTino "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/adapters/adapter_tino"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	sharedDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/shared"
)

type TinoAdapter struct {
	client adapterTino.TinoService
}

func NewTinoAdapter(input sharedDTO.TinoAdapterDTO) *TinoAdapter {
	return &TinoAdapter{client: adapterTino.New(adapterTino.InitInput{
		Url:         input.Url,
		User:        input.User,
		Password:    input.Password,
		CallbackUrl: input.CallbackUrl,
	})}
}

func (a *TinoAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("tino adapter not configured")
	}

	res, err := a.client.CreateInvoice(payment)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: res.InvoiceID,
		IsPaid:        res.Status == adapterTino.INVOICE_STATUS_PAID,
		Raw:           res,
	}, nil
}

func (a *TinoAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("tino adapter not configured")
	}

	err := a.client.CheckInvoice(payment.RefInvoiceID)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: true,
	}, nil
}
