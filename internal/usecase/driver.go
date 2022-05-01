package usecase

import (
	"github.com/febriliankr/go-clean-architecture/internal"
	"github.com/febriliankr/go-clean-architecture/internal/entity"
)

type DriverUsecase struct {
	repo internal.DriverRepo
}

func NewDriverUsecase(repo internal.DriverRepo) *DriverUsecase {
	return &DriverUsecase{
		repo: repo,
	}
}

func (uc *DriverUsecase) CreateNewDriver(req entity.CreateNewDriverRequest) (entity.CreateNewDriverResponse, error) {
	return uc.repo.CreateNewDriver(req)
}

func (uc *DriverUsecase) GetDriverList(req entity.GetDriverListRequest) (entity.GetDriverListResponse, error) {
	return uc.repo.GetDriverList(req)
}

func (uc *DriverUsecase) GetDriverByID(req entity.GetDriverByIDRequest) (entity.GetDriverByIDResponse, error) {
	return uc.repo.GetDriverByID(req)
}

func (uc *DriverUsecase) UpdateDriver(req entity.UpdateDriverRequest) (entity.UpdateDriverResponse, error) {
	return uc.repo.UpdateDriver(req)
}

func (uc *DriverUsecase) DeleteDriver(req entity.DeleteDriverRequest) (entity.DeleteDriverResponse, error) {
	return uc.repo.DeleteDriver(req)
}
