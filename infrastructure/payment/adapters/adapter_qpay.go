package adapters

import (
	"fmt"
	"strconv"

	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	paymentServiceResponseDTO "github.com/techpartners-asia/payments-service/infrastructure/payment/dto/response"
	sharedDTO "github.com/techpartners-asia/payments-service/infrastructure/shared"

	"github.com/techpartners-asia/qpay-go/qpay_v2"
)

// QPayAdapter implements PaymentProvider for QPay.
type QPayAdapter struct {
	client qpay_v2.QPay
}

func NewQPayAdapter(input sharedDTO.QpayAdapterDTO) *QPayAdapter {
	if input.Username == "" || input.Password == "" || input.Endpoint == "" || input.InvoiceCode == "" || input.MerchantID == "" {
		return nil
	}
	return &QPayAdapter{client: qpay_v2.New(input.Username, input.Password, input.Endpoint, input.Callback, input.InvoiceCode, input.MerchantID)}
}

func (a *QPayAdapter) CreateInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("qpay adapter not configured")
	}
	// prefix := "personal"
	// if input.IsOrg && input.OrgRegNo != "" {
	// 	prefix = input.OrgRegNo
	// }
	qpayInput := qpay_v2.QPayCreateInvoiceInput{
		SenderCode:    payment.UID,
		ReceiverCode:  payment.UID,
		Description:   payment.Note,
		Amount:        int64(payment.Amount),
		CallbackParam: map[string]string{"uid": payment.UID},
	}
	res, _, err := a.client.CreateInvoice(qpayInput)
	if err != nil {
		return nil, err
	}

	var deeplinks []paymentServiceResponseDTO.Deeplink
	for _, v := range res.Urls {
		deeplinks = append(deeplinks, paymentServiceResponseDTO.Deeplink{
			Name:        v.Name,
			Description: v.Description,
			Link:        v.Link,
			Logo:        v.Logo,
		})
	}

	return &paymentServiceResponseDTO.InvoiceResult{
		BankInvoiceID: res.InvoiceID,
		BankQRCode:    res.QrText,
		Deeplinks:     deeplinks,
		IsPaid:        false,
		Raw:           res,
	}, nil
}

func (a *QPayAdapter) CheckInvoice(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("qpay adapter not configured")
	}

	res, _, err := a.client.CheckPayment(payment.UID, 100, 1)
	if err != nil {
		return nil, err
	}

	amount := float64(0)
	for _, row := range res.Rows {
		if row.PaymentStatus == "PAID" {
			if _amount, err := strconv.ParseFloat(row.PaymentAmount, 64); err == nil {
				amount += _amount
			}
		}
	}

	return &paymentServiceResponseDTO.CheckInvoiceResult{
		IsPaid: amount >= payment.Amount,
	}, nil
}
