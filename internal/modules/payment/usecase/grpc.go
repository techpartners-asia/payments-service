package usecasePayment

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	paymentService "github.com/techpartners-asia/payments-service/infrastructure/payment"
	repositoryRedis "github.com/techpartners-asia/payments-service/infrastructure/redis/repository"
	grpcMapDTO "github.com/techpartners-asia/payments-service/internal/modules/payment/dto/map/grpc"
	grpcRequestDTO "github.com/techpartners-asia/payments-service/internal/modules/payment/dto/request/grpc"
	grpcResponseDTO "github.com/techpartners-asia/payments-service/internal/modules/payment/dto/response/grpc"
	repositoryPayment "github.com/techpartners-asia/payments-service/internal/modules/payment/repository"
	paymentProto "github.com/techpartners-asia/payments-service/pkg/proto/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type paymentUsecase struct {
	paymentRepo     repositoryPayment.PaymentRepository
	redisRepository repositoryRedis.RedisRepository

	paymentProto.UnimplementedPaymentServiceServer
}

func NewPaymentUsecase(paymentRepo repositoryPayment.PaymentRepository, redisRepository repositoryRedis.RedisRepository) paymentProto.PaymentServiceServer {
	return &paymentUsecase{
		paymentRepo:     paymentRepo,
		redisRepository: redisRepository,
	}
}

func (u *paymentUsecase) Create(ctx context.Context, req *paymentProto.PaymentCreateRequest) (*paymentProto.PaymentCreateResponse, error) {

	payment, err := u.paymentRepo.CreatePayment(grpcRequestDTO.ToEntity(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// * : Add payment service to create payment
	result, err := paymentService.NewPaymentService(payment.MerchantID, u.paymentRepo, u.redisRepository).Create(payment)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// * : Update payment invoice id
	updatedInstance, err := u.paymentRepo.UpdateInvoiceID(payment.UID, result.BankInvoiceID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return grpcResponseDTO.ToCreateResponse(updatedInstance), nil
}

func (u *paymentUsecase) Check(ctx context.Context, req *paymentProto.PaymentCheckRequest) (*paymentProto.PaymentCheckResponse, error) {
	payment, err := u.paymentRepo.GetByUID(req.Uid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// * : Add payment service to check payment
	result, err := paymentService.NewPaymentService(payment.MerchantID, u.paymentRepo, u.redisRepository).Check(payment)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if result.IsPaid {
		// * : Update payment status
		updatedInstance, err := u.paymentRepo.UpdatePaymentStatus(payment.UID, entity.PaymentStatusPaid)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		return grpcResponseDTO.ToCheckResponse(updatedInstance), nil
	}

	return grpcResponseDTO.ToCheckResponse(payment), nil
}

func (u *paymentUsecase) CheckStream(req *paymentProto.PaymentCheckRequest, stream paymentProto.PaymentService_CheckStreamServer) error {
	payment, err := u.paymentRepo.GetByUID(req.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Errorf(codes.NotFound, "Payment with UID %s not found", req.Uid)
		}
		return status.Errorf(codes.Internal, "Database error: %v", err)
	}

	initialResponse := &paymentProto.PaymentCheckResponse{
		Uid:    payment.UID,
		Status: grpcMapDTO.ToPaymentStatus(payment.Status),
	}
	if err := stream.Send(initialResponse); err != nil {
		return status.Errorf(codes.Unavailable, "Failed to send initial response: %v", err)
	}

	for payment.Status == entity.PaymentStatusPending {

		select {
		case <-time.After(5 * time.Second):
		case <-stream.Context().Done():
			log.Printf("Client disconnected during streaming for UID: %s", req.Uid)
			return nil
		}

		payment, err = u.paymentRepo.GetByUID(req.Uid)
		if err != nil {
			return status.Errorf(codes.Internal, "Error re-fetching payment status: %v", err)
		}

		// * : Add payment service to check payment
		result, err := paymentService.NewPaymentService(payment.MerchantID, u.paymentRepo, u.redisRepository).Check(payment)
		if err != nil {
			return status.Errorf(codes.Internal, "Error checking payment: %v", err)
		}

		if result.IsPaid {
			// * : Update payment status
			updatedInstance, err := u.paymentRepo.UpdatePaymentStatus(payment.UID, entity.PaymentStatusPaid)
			if err != nil {
				return status.Errorf(codes.Internal, "Error updating payment status: %v", err)
			}

			if err := stream.Send(grpcResponseDTO.ToCheckResponse(updatedInstance)); err != nil {
				return status.Errorf(codes.Internal, "Error sending updated payment status: %v", err)
			}

			return nil
		} else {
			if err := stream.Send(grpcResponseDTO.ToCheckResponse(payment)); err != nil {
				return status.Errorf(codes.Internal, "Error sending updated payment status: %v", err)
			}
		}
	}

	log.Printf("Streaming completed for UID: %s. Final status: %s", req.GetUid(), payment.Status)

	return nil
}
