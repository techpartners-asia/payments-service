package adapters

import (
	"fmt"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	sharedDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/shared"

	tokipay "github.com/techpartners-asia/tokipay-go"
)

// TokiPayAdapter implements PaymentProvider for Tokipay.
type TokiPayAdapter struct {
	client tokipay.Tokipay
}

func NewTokiPayAdapter(input sharedDTO.TokipayAdapterDTO) *TokiPayAdapter {
	return &TokiPayAdapter{client: tokipay.New(input.Endpoint, input.APIKey, input.IMAPIKey, input.Authorization, input.MerchantID, input.SuccessURL, input.FailureURL, input.AppSchemaIOS)}
}

func (a *TokiPayAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("tokipay adapter not configured")
	}
	res, err := a.client.PaymentSentUser(tokipay.TokipayPaymentInput{
		OrderId:     payment.UID,
		Amount:      int64(payment.Amount),
		PhoneNo:     payment.Phone,
		CountryCode: "+976",
		Notes:       payment.Note,
	})
	if err != nil {
		return nil, err
	}
	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: payment.UID,
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *TokiPayAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("tokipay adapter not configured")
	}

	res, err := a.client.PaymentStatus(payment.UID)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: res.Data.Status == "COMPLETED",
	}, nil
}
