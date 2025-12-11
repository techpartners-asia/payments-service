package grpcMapDTO

import (
	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	paymentProto "github.com/techpartners-asia/payments-service/pkg/proto/payment"
)

func ToPaymentStatus(status entity.PaymentStatus) paymentProto.PaymentStatus {
	switch status {
	case entity.PaymentStatusPending:
		return paymentProto.PaymentStatus_PAYMENT_STATUS_PENDING
	case entity.PaymentStatusPaid:
		return paymentProto.PaymentStatus_PAYMENT_STATUS_PAID
	case entity.PaymentStatusCancelled:
		return paymentProto.PaymentStatus_PAYMENT_STATUS_CANCELLED
	case entity.PaymentStatusRefunded:
		return paymentProto.PaymentStatus_PAYMENT_STATUS_REFUNDED
	default:
		return paymentProto.PaymentStatus_PAYMENT_STATUS_PENDING
	}
}
