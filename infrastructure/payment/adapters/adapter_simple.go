package adapters

import (
	"fmt"
	"time"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	sharedDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/shared"

	simple "github.com/techpartners-asia/simple-go"
)

// SimpleAdapter implements PaymentProvider for Simple.
type SimpleAdapter struct {
	client   simple.Simple
	simpleID string
}

func NewSimpleAdapter(input sharedDTO.SimpleAdapterDTO) *SimpleAdapter {
	return &SimpleAdapter{client: simple.New(input.UserName, input.Password, input.BaseUrl, input.CallbackUrl), simpleID: input.SimpleID}
}

func (a *SimpleAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("simple adapter not configured")
	}

	expireAt := time.Now().Add(20 * time.Minute).Format("2006-01-02 15:04:05")

	req := simple.SimpleCreateInvoiceInput{
		OrderID:    payment.UID,
		Total:      int(payment.Amount),
		ExpireDate: expireAt,
	}

	res, err := a.client.CreateInvoice(req)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: payment.UID,
		Raw:           res,
		IsPaid:        false,
	}, nil
}

func (a *SimpleAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("simple adapter not configured")
	}

	res, err := a.client.GetInvoice(simple.SimpleGetInvoiceRequest{
		OrderID:  payment.UID,
		SimpleID: "",
	})
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: res.Data.InvoiceStatus == "PAID",
	}, nil
}
