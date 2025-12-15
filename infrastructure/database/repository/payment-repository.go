package repositoryImpl

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	"gorm.io/gorm"
)

type paymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepositoryImpl {
	return &paymentRepositoryImpl{db: db}
}

func (r *paymentRepositoryImpl) CreatePayment(payment *entity.PaymentEntity) (*entity.PaymentEntity, error) {
	if err := r.db.Create(payment).Error; err != nil {
		return nil, err
	}
	return r.GetPaymentByID(payment.ID)
}

func (r *paymentRepositoryImpl) GetPaymentByID(id uint) (*entity.PaymentEntity, error) {
	var payment entity.PaymentEntity
	if err := r.db.Where("id = ?", id).Preload("Merchant").First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepositoryImpl) GetByUID(uid string) (*entity.PaymentEntity, error) {
	var payment entity.PaymentEntity
	if err := r.db.Where("uid = ?", uid).Preload("Merchant").First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepositoryImpl) UpdatePayment(payment *entity.PaymentEntity) (*entity.PaymentEntity, error) {
	if err := r.db.Save(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *paymentRepositoryImpl) DeletePayment(id uint) error {
	if err := r.db.Delete(&entity.PaymentEntity{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *paymentRepositoryImpl) UpdateInvoiceID(uid string, invoiceID string) (*entity.PaymentEntity, error) {
	var payment entity.PaymentEntity
	if err := r.db.Where("uid = ?", uid).First(&payment).Error; err != nil {
		return nil, err
	}
	payment.RefInvoiceID = invoiceID
	if err := r.db.Save(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepositoryImpl) UpdatePaymentStatus(uid string, status entity.PaymentStatus) (*entity.PaymentEntity, error) {
	var payment entity.PaymentEntity
	if err := r.db.Where("uid = ?", uid).First(&payment).Error; err != nil {
		return nil, err
	}
	payment.Status = status
	if err := r.db.Save(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}
