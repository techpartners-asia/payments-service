package grpcResponseDTO

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	paymentServiceResponseDTO "git.techpartners.asia/gateway-services/payment-service/infrastructure/payment/dto/response"
	paymentProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/payment"
)

// ToCreateResponse converts an internal Entity into a Protobuf response.
func ToCreateResponse(p *entity.PaymentEntity, result *paymentServiceResponseDTO.InvoiceResult) *paymentProto.PaymentCreateResponse {
	return &paymentProto.PaymentCreateResponse{
		Uid: p.UID,
		InvoiceResult: &paymentProto.InvoiceResult{
			BankInvoiceId: result.BankInvoiceID,
			BankQrCode:    result.BankQRCode,
			Deeplinks:     ToDeeplinks(result.Deeplinks),
			IsPaid:        result.IsPaid,
			// Raw:           anypb.New(&result.Raw),
		},
	}
}

func ToDeeplinks(deeplinks []paymentServiceResponseDTO.Deeplink) []*paymentProto.Deeplink {
	var deeplinksProto []*paymentProto.Deeplink
	for _, deeplink := range deeplinks {
		deeplinksProto = append(deeplinksProto, &paymentProto.Deeplink{
			Name:        deeplink.Name,
			Description: deeplink.Description,
			Link:        deeplink.Link,
			Logo:        deeplink.Logo,
		})
	}
	return deeplinksProto
}
