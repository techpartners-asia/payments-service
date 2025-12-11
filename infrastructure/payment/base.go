package paymentService

import (
	"errors"

	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	"github.com/techpartners-asia/payments-service/infrastructure/payment/adapters"
	paymentServiceResponseDTO "github.com/techpartners-asia/payments-service/infrastructure/payment/dto/response"
	repositoryRedis "github.com/techpartners-asia/payments-service/infrastructure/redis/repository"
	usecaseRedis "github.com/techpartners-asia/payments-service/infrastructure/redis/usecase"
	sharedDTO "github.com/techpartners-asia/payments-service/infrastructure/shared"
	repositoryPayment "github.com/techpartners-asia/payments-service/internal/modules/payment/repository"
)

type PaymentService interface {
	Create(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error)
	Check(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error)
}

type paymentService struct {
	paymentRepo repositoryPayment.PaymentRepository

	configs sharedDTO.SharedPaymentConfigDTO
}

func NewPaymentService(merchantID uint, paymentRepo repositoryPayment.PaymentRepository, redisRepository repositoryRedis.RedisRepository) PaymentService {

	merchantUsecase := usecaseRedis.NewMerchantUsecase(redisRepository)
	merchant, err := merchantUsecase.Get(merchantID)
	if err != nil {
		panic(err)
	}

	return &paymentService{
		paymentRepo: paymentRepo,
		configs:     merchant.Configs,
	}
}

func (s *paymentService) Create(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.InvoiceResult, error) {

	switch payment.Type {
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
		// case entity.PaymentTypeBpay:
		// 	return adapters.NewBpayAdapter(s.configs.Bpay).CreateInvoice(payment)
		// case entity.PaymentTypeEcommerce:
		// 	return adapters.NewEcommerceAdapter(s.configs.Ecommerce).CreateInvoice(payment)
		// case entity.PaymentTypeHipay:
		// 	return adapters.NewHipayAdapter(s.configs.Hipay).CreateInvoice(payment)
		// case entity.PaymentTypeMongolchat:
		// 	return adapters.NewMongolchatAdapter(s.configs.Mongolchat).CreateInvoice(payment)
		// case entity.PaymentTypePass:
		// 	return adapters.NewPassAdapter(s.configs.Pass).CreateInvoice(payment)
	}
	return nil, errors.New("invalid payment type")
}

func (s *paymentService) Check(payment *entity.PaymentEntity) (*paymentServiceResponseDTO.CheckInvoiceResult, error) {

	switch payment.Type {
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
