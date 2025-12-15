package paymentServiceRequestDTO

import "git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"

type (
	InvoiceInput struct {
		Amount      float64            // Amount
		UID         string             // payment uid or order uid
		Phone       string             // phone number
		CustomerID  uint               // customer id
		Note        string             // Note : description of the invoice
		CallbackURL string             // CallbackURL : callback url
		ReturnType  string             // ReturnType : return type
		Type        entity.PaymentType // qpay , tokipay , monpay , golomt , socialpay , storepay , pocket , simple , balc
	}

	CheckInvoiceInput struct {
		UID    string             `json:"uid"`
		Amount float64            `json:"amount"`
		Type   entity.PaymentType `json:"type"`
	}
)
