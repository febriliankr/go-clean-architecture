package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type ShipmentDB struct {
	db *sqlx.DB
}

func NewShipmentDB(db *sqlx.DB) *ShipmentDB {
	return &ShipmentDB{
		db: db,
	}
}

func (r *ShipmentDB) CreateNewShipment(req entity.CreateNewShipmentRequest) (entity.CreateNewShipmentResponse, error) {
	var res entity.CreateNewShipmentResponse
	q, err := r.db.Prepare(queryCreateNewShipment)
	if err != nil {
		return res, err
	}

	err = q.QueryRow(req.ShipmentNumber, req.Origin, req.Destination, req.LoadingDate).Scan(&res.ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *ShipmentDB) AllocateShipment(req entity.AllocateShipmentRequest) (entity.AllocateShipmentResponse, error) {
	var res entity.AllocateShipmentResponse
	return res, nil
}

func (r *ShipmentDB) GetShipmentList(req entity.GetShipmentListRequest) (entity.GetShipmentListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ShipmentDB) GetShipmentByID(req entity.GetShipmentByIDRequest) (entity.GetShipmentByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ShipmentDB) UpdateShipment(req entity.UpdateShipmentRequest) (entity.UpdateShipmentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ShipmentDB) DeleteShipment(req entity.DeleteShipmentRequest) (entity.DeleteShipmentResponse, error) {
	//TODO implement me
	panic("implement me")
}
