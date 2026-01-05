package adapterTino

import (
	"errors"
	"fmt"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	"github.com/go-resty/resty/v2"
)

type TinoService interface {
	CreateInvoice(payment *entity.PaymentEntity) (*CreateInvoiceDataResponse, error)
	CheckInvoice(invoiceID string) error
	CancelInvoice(invoiceID string) error
	RefundInvoice(invoiceID string, reason string) error
}

type tinoService struct {
	resty.Client
	input InitInput
}

func New(input InitInput) TinoService {

	auth := auth(input)

	if !auth.Status {
		return nil
	}

	return &tinoService{
		Client: *resty.New().SetBaseURL(input.Url).SetHeader("Authorization", fmt.Sprintf("Bearer %v", auth.Data.Token)),
		input:  input,
	}
}

func auth(input InitInput) AuthResponse {
	var response AuthResponse
	if _, err := resty.New().R().
		SetBody(map[string]interface{}{
			"username": input.User,
			"password": input.Password,
		}).
		SetResult(&response).
		Post(fmt.Sprintf("%v/login", input.Url)); err != nil {

		return AuthResponse{}
	}

	return response
}

func (s *tinoService) RefundInvoice(invoiceID string, reason string) error {

	var response RefundInvoiceResponse

	if _, err := s.Client.R().SetBody(
		&RefundInvoiceInput{
			InvoiceID: invoiceID,
			Reason:    reason,
		},
	).SetResult(&response).Post("/invoice/refund"); err != nil {
		return err
	}

	if !response.Status {
		return errors.New(response.Message)
	}

	return nil
}

func (s *tinoService) CreateInvoice(payment *entity.PaymentEntity) (*CreateInvoiceDataResponse, error) {

	var response CreateInvoiceResponse

	if _, err := s.Client.R().SetBody(CreateInvoiceInput{
		CallbackURL:   fmt.Sprintf("%v/%v", s.input.CallbackUrl, payment.UID),
		Description:   payment.Note,
		Amount:        payment.Amount,
		UserID:        payment.CustomerID,
		TransactionID: payment.UID,
		// WalletID:      walletID,
	}).SetResult(&response).Post("/invoice"); err != nil {
		return nil, err
	}

	if !response.Status {
		return nil, errors.New(response.Message)
	}

	return &response.Data, nil
}

func (s *tinoService) CheckInvoice(invoiceID string) error {

	var response CheckInvoiceResponse

	if _, err := s.Client.R().SetResult(&response).Get(fmt.Sprintf("/invoice/check/%v", invoiceID)); err != nil {
		return err
	}

	if !response.Status {
		return errors.New(response.Message)
	}

	if response.Data.Status == INVOICE_STATUS_PAID {
		return nil
	}

	return errors.New("Нэхэмжлэх төлөгдөөгүй байна")
}

func (s *tinoService) CancelInvoice(invoiceID string) error {

	var response CancelInvoiceResponse

	if _, err := s.Client.R().SetResult(&response).Get(fmt.Sprintf("/invoice/cancel/%v", invoiceID)); err != nil {
		return err
	}

	if !response.Status {
		return errors.New(response.Message)
	}

	return nil
}
