package entity

import (
	databaseBase "git.techpartners.asia/gateway-services/payment-service/infrastructure/database/base"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusPaid      PaymentStatus = "paid"
	PaymentStatusCancelled PaymentStatus = "cancelled"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

type (
	PaymentEntity struct {
		databaseBase.BaseEntity
		UID          string             `gorm:"uid;index" json:"uid"`
		Status       PaymentStatus      `gorm:"status" json:"status"`
		Amount       float64            `gorm:"amount" json:"amount"`
		Phone        string             `gorm:"phone" json:"phone"`
		CustomerID   uint               `gorm:"customer_id;index" json:"customer_id"`
		Note         string             `gorm:"note" json:"note"`
		RefInvoiceID string             `gorm:"ref_invoice_id" json:"ref_invoice_id"`
		MerchantID   uint               `gorm:"merchant_id;index" json:"merchant_id"`
		Merchant     *MerchantEntity    `gorm:"foreignKey:MerchantID" json:"merchant"`
		PaymentType  PaymentType        `gorm:"payment_type" json:"payment_type"`
		Logs         []PaymentLogEntity `gorm:"foreignKey:PaymentID" json:"logs"`
	}

	PaymentLogEntity struct {
		databaseBase.BaseEntity
		PaymentID uint           `gorm:"payment_id;index" json:"payment_id"`
		Payment   *PaymentEntity `gorm:"foreignKey:PaymentID" json:"payment"`
		Message   string         `gorm:"message" json:"message"`
	}
)

func (p *PaymentEntity) TableName() string {
	return "payments"
}
func (p *PaymentLogEntity) TableName() string {
	return "payment_logs"
}
