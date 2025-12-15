package adapters

import (
	"fmt"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	sharedDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/shared"

	storepay "github.com/techpartners-asia/storepay-go"
)

// StorePayAdapter implements PaymentProvider for StorePay.
type StorePayAdapter struct {
	client storepay.Storepay
}

func NewStorePayAdapter(input sharedDTO.StorePayAdapterDTO) *StorePayAdapter {
	return &StorePayAdapter{client: storepay.New(input.AppUserName, input.AppPassword, input.Username, input.Password, input.AuthUrl, input.BaseUrl, input.StoreId, input.CallbackUrl)}
}

func (a *StorePayAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("storepay adapter not configured")
	}

	res, err := a.client.Loan(storepay.StorepayLoanInput{
		Amount:       payment.Amount,
		MobileNumber: payment.Phone,
		Description:  payment.Note,
	})
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: fmt.Sprintf("%d", res),
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *StorePayAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("storepay adapter not configured")
	}

	res, err := a.client.LoanCheck(payment.UID)
	if err != nil {
		return nil, err
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: res,
	}, nil
}
