package paymentService

import (
	"errors"

	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/adapters"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	repositoryRedis "git.techpartners.asia/gateway-services/payment-service/infrastructure/redis/repository"
	usecaseRedis "git.techpartners.asia/gateway-services/payment-service/infrastructure/redis/usecase"
	sharedDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/shared"
)

type PaymentService interface {
	Create(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error)
	Check(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error)
}

type paymentService struct {
	configs sharedDTO.SharedPaymentConfigDTO
}

func NewPaymentService(merchantUID string, redisRepository repositoryRedis.RedisRepository) PaymentService {

	merchantUsecase := usecaseRedis.NewMerchantUsecase(redisRepository)
	merchant, err := merchantUsecase.Get(merchantUID)
	if err != nil {
		panic(err)
	}

	return &paymentService{
		configs: merchant.Configs,
	}
}

func (s *paymentService) Create(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {

	switch payment.PaymentType {
	case entity.PaymentTypeQpay:
		return adapters.NewQPayAdapter(s.configs.Qpay).CreateInvoice(payment)
	case entity.PaymentTypeTokipay:
		return adapters.NewTokiPayAdapter(s.configs.Tokiay).CreateInvoice(payment)
	case entity.PaymentTypeMonpay:
		return adapters.NewMonpayAdapter(s.configs.Monpay).CreateInvoice(payment)
	case entity.PaymentTypeEcommerce:
		return adapters.NewGolomtAdapter(s.configs.Golomt).CreateInvoice(payment)
	case entity.PaymentTypeSocialpay:
		return adapters.NewSocialPayAdapter(s.configs.Socialpay).CreateInvoice(payment)
	case entity.PaymentTypeStorepay:
		return adapters.NewStorePayAdapter(s.configs.Storepay).CreateInvoice(payment)
	case entity.PaymentTypePocket:
		return adapters.NewPocketAdapter(s.configs.Pocket).CreateInvoice(payment)
	case entity.PaymentTypeSimple:
		return adapters.NewSimpleAdapter(s.configs.Simple).CreateInvoice(payment)
	case entity.PaymentTypeBalc:
		return adapters.NewBalcCreditAdapter(s.configs.Balc).CreateInvoice(payment)
	case entity.PaymentTypeTino:
		return adapters.NewTinoAdapter(s.configs.Tino).CreateInvoice(payment)
	case entity.PaymentTypeBpay:
		return adapters.NewBpayAdapter(s.configs.Bpay).CreateInvoice(payment)
	case entity.PaymentTypeHipay:
		return adapters.NewHipayAdapter(s.configs.Hipay).CreateInvoice(payment)
		// case entity.PaymentTypeMongolchat:
		// 	return adapters.NewMongolchatAdapter(s.configs.Mongolchat).CreateInvoice(payment)
		// case entity.PaymentTypePass:
		// 	return adapters.NewPassAdapter(s.configs.Pass).CreateInvoice(payment)
	}
	return nil, errors.New("invalid payment type")
}

func (s *paymentService) Check(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {

	switch payment.PaymentType {
	case entity.PaymentTypeQpay:
		return adapters.NewQPayAdapter(s.configs.Qpay).CheckInvoice(payment)
	case entity.PaymentTypeTokipay:
		return adapters.NewTokiPayAdapter(s.configs.Tokiay).CheckInvoice(payment)
	case entity.PaymentTypeMonpay:
		return adapters.NewMonpayAdapter(s.configs.Monpay).CheckInvoice(payment)
	case entity.PaymentTypeEcommerce:
		return adapters.NewGolomtAdapter(s.configs.Golomt).CheckInvoice(payment)
	case entity.PaymentTypeSocialpay:
		return adapters.NewSocialPayAdapter(s.configs.Socialpay).CheckInvoice(payment)
	case entity.PaymentTypeStorepay:
		return adapters.NewStorePayAdapter(s.configs.Storepay).CheckInvoice(payment)
	case entity.PaymentTypePocket:
		return adapters.NewPocketAdapter(s.configs.Pocket).CheckInvoice(payment)
	case entity.PaymentTypeSimple:
		return adapters.NewSimpleAdapter(s.configs.Simple).CheckInvoice(payment)
	case entity.PaymentTypeBalc:
		return adapters.NewBalcCreditAdapter(s.configs.Balc).CheckInvoice(payment)
		// case entity.PaymentTypeBpay:
		// 	return adapters.NewBpayAdapter(s.configs.Bpay).CheckInvoice(payment)
		// case entity.PaymentTypeEcommerce:
		// 	return adapters.NewEcommerceAdapter(s.configs.Ecommerce).CheckInvoice(payment)
		// case entity.PaymentTypeHipay:
		// 	return adapters.NewHipayAdapter(s.configs.Hipay).CheckInvoice(payment)
		// case entity.PaymentTypeMongolchat:
		// 	return adapters.NewMongolchatAdapter(s.configs.Mongolchat).CheckInvoice(payment)
		// case entity.PaymentTypePass:
		// 	return adapters.NewPassAdapter(s.configs.Pass).CheckInvoice(payment)
	}
	return nil, errors.New("invalid payment type")
}
