package repositoryPayment

import "github.com/techpartners-asia/payments-service/infrastructure/database/entity"

type PaymentRepository interface {
	CreatePayment(payment *entity.PaymentEntity) (*entity.PaymentEntity, error)
	GetPaymentByID(id uint) (*entity.PaymentEntity, error)
	GetByUID(uid string) (*entity.PaymentEntity, error)
	UpdatePayment(payment *entity.PaymentEntity) (*entity.PaymentEntity, error)
	DeletePayment(id uint) error
	UpdateInvoiceID(uid string, invoiceID string) (*entity.PaymentEntity, error)
	UpdatePaymentStatus(uid string, status entity.PaymentStatus) (*entity.PaymentEntity, error)
}
