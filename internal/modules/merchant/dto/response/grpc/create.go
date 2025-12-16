package grpcMerchantResponseDTO

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	merchantProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/merchant"
)

func ToResponse(merchant *entity.MerchantEntity) *merchantProto.MerchantResponse {
	return &merchantProto.MerchantResponse{
		Uid:  merchant.UID,
		Name: merchant.Name,
	}
}
