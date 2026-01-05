package grpcRequestDTO

import (
	"strconv"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"

	paymentProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/payment"
)

func ToEntity(req *paymentProto.PaymentCreateRequest) *entity.PaymentEntity {
	return &entity.PaymentEntity{
		Amount:      float64(req.Amount),
		Status:      entity.PaymentStatusPending,
		PaymentType: ToType(req.Type),
		MerchantID:  uint(req.MerchantId),
		Phone:       req.Phone,
		CustomerID:  strconv.FormatUint(req.CustomerId, 10),
		Note:        req.Note,
	}
}

func ToType(typeProto paymentProto.PaymentType) entity.PaymentType {
	switch typeProto {
	case paymentProto.PaymentType_PAYMENT_TYPE_QPAY:
		return entity.PaymentTypeQpay
	case paymentProto.PaymentType_PAYMENT_TYPE_TOKIPAY:
		return entity.PaymentTypeTokipay
	case paymentProto.PaymentType_PAYMENT_TYPE_MONPAY:
		return entity.PaymentTypeMonpay
	case paymentProto.PaymentType_PAYMENT_TYPE_GOLOMT:
		return entity.PaymentTypeEcommerce
	case paymentProto.PaymentType_PAYMENT_TYPE_SOCIALPAY:
		return entity.PaymentTypeSocialpay
	case paymentProto.PaymentType_PAYMENT_TYPE_STOREPAY:
		return entity.PaymentTypeStorepay
	case paymentProto.PaymentType_PAYMENT_TYPE_POCKET:
		return entity.PaymentTypePocket
	case paymentProto.PaymentType_PAYMENT_TYPE_BALC:
		return entity.PaymentTypeBalc
	case paymentProto.PaymentType_PAYMENT_TYPE_SIMPLE:
		return entity.PaymentTypeSimple
	default:
		return entity.PaymentTypeSimple
	}
}
