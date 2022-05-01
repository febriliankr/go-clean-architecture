package usecase

import (
	"math/rand"
	"time"

	"github.com/febriliankr/go-clean-architecture/internal"
	"github.com/febriliankr/go-clean-architecture/internal/entity"
)

type ShipmentUsecase struct {
	repo internal.ShipmentRepo
}

func NewShipmentUsecase(repo internal.ShipmentRepo) *ShipmentUsecase {
	return &ShipmentUsecase{
		repo: repo,
	}
}

func (uc *ShipmentUsecase) CreateNewShipment(req entity.CreateNewShipmentRequest) (entity.CreateNewShipmentResponse, error) {
	req.ShipmentNumber = randSeq(8)
	parse, err := time.Parse("2006-01-02T15:04:05.000Z", req.LoadingDateStr)
	if err != nil {
		return entity.CreateNewShipmentResponse{}, err
	}
	req.LoadingDate = parse

	return uc.repo.CreateNewShipment(req)
}

func (uc *ShipmentUsecase) AllocateShipment(req entity.AllocateShipmentRequest) (entity.AllocateShipmentResponse, error) {
	return uc.repo.AllocateShipment(req)
}

func (uc *ShipmentUsecase) GetShipmentList(req entity.GetShipmentListRequest) (entity.GetShipmentListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *ShipmentUsecase) GetShipmentByID(req entity.GetShipmentByIDRequest) (entity.GetShipmentByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *ShipmentUsecase) UpdateShipment(req entity.UpdateShipmentRequest) (entity.UpdateShipmentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *ShipmentUsecase) DeleteShipment(req entity.DeleteShipmentRequest) (entity.DeleteShipmentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
