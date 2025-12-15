package adapters

import (
	"fmt"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	sharedDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/shared"

	balcapi "github.com/techpartners-asia/balc-api-go"
)

// BalcCreditAdapter implements PaymentProvider for Balc credit flow.
type BalcCreditAdapter struct {
	client balcapi.Balc
}

func NewBalcCreditAdapter(input sharedDTO.BalcAdapterDTO) *BalcCreditAdapter {
	return &BalcCreditAdapter{client: balcapi.New(input.Endpoint, input.Token)}
}

func (a *BalcCreditAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("balc adapter not configured")
	}

	creditCheck, err := a.client.LimitCheck(int(payment.CustomerID))
	if err != nil {
		return nil, fmt.Errorf("error on balcAPI check: %w", err)
	}
	if creditCheck.AvailLimit < payment.Amount {
		return nil, fmt.Errorf("таны кредит гүйлгээний дүнд хүрэхгүй байна")
	}

	loanAccountID, err := a.client.Loan(int(payment.Amount), "Зээл", int(payment.CustomerID))
	if err != nil {
		return nil, fmt.Errorf("зээл авахад алдаа гарлаа: %w", err)
	}

	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: fmt.Sprintf("%v", loanAccountID),
		IsPaid:        true,
		Raw:           loanAccountID,
	}, nil
}

func (a *BalcCreditAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("balc adapter not configured")
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: true,
	}, nil
}
