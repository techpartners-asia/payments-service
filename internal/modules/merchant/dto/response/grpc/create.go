package grpcMerchantResponseDTO

import (
	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	merchantProto "github.com/techpartners-asia/payments-service/pkg/proto/merchant"
)

func ToResponse(merchant *entity.MerchantEntity) *merchantProto.MerchantResponse {
	return &merchantProto.MerchantResponse{
		Id:   uint64(merchant.ID),
		Name: merchant.Name,
	}
}
