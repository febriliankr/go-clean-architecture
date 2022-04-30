package usecase

import (
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type TruckUsecase struct {
	repo internal.TruckRepo
}

func NewTruckUsecase(repo internal.TruckRepo) *TruckUsecase {
	return &TruckUsecase{
		repo: repo,
	}
}

func (uc *TruckUsecase) CreateNewTruck(req entity.CreateNewTruckRequest) (entity.CreateNewTruckResponse, error) {
	return uc.repo.CreateNewTruck(req)
}

func (uc *TruckUsecase) GetTruckList(req entity.GetTruckListRequest) (entity.GetTruckListResponse, error) {
	return uc.repo.GetTruckList(req)
}

func (uc *TruckUsecase) GetTruckByID(req entity.GetTruckByIDRequest) (entity.GetTruckByIDResponse, error) {
	return uc.repo.GetTruckByID(req)
}

func (uc *TruckUsecase) UpdateTruck(req entity.UpdateTruckRequest) (entity.UpdateTruckResponse, error) {
	return uc.repo.UpdateTruck(req)
}

func (uc *TruckUsecase) DeleteTruck(req entity.DeleteTruckRequest) (entity.DeleteTruckResponse, error) {
	return uc.repo.DeleteTruck(req)
}
