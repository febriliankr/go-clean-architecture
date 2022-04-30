package usecase

import (
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type PaymentUsecase struct {
	repo internal.PaymentRepo
}

func NewPaymentUsecase(repo internal.PaymentRepo) *PaymentUsecase {
	return &PaymentUsecase{
		repo: repo,
	}
}

func (uc *PaymentUsecase) CreateNewPayment(req entity.CreateNewPaymentRequest) (entity.CreateNewPaymentResponse, error) {
	return uc.repo.CreateNewPayment(req)
}

func (uc *PaymentUsecase) GetPaymentList(req entity.GetPaymentListRequest) (entity.GetPaymentListResponse, error) {
	return uc.repo.GetPaymentList(req)
}

func (uc *PaymentUsecase) GetPaymentByID(req entity.GetPaymentByIDRequest) (entity.GetPaymentByIDResponse, error) {
	return uc.repo.GetPaymentByID(req)
}

func (uc *PaymentUsecase) UpdatePayment(req entity.UpdatePaymentRequest) (entity.UpdatePaymentResponse, error) {
	return uc.repo.UpdatePayment(req)
}

func (uc *PaymentUsecase) DeletePayment(req entity.DeletePaymentRequest) (entity.DeletePaymentResponse, error) {
	return uc.repo.DeletePayment(req)
}
