package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type TruckDB struct {
	db *sqlx.DB
}

func NewTruckDB(db *sqlx.DB) *TruckDB {
	return &TruckDB{
		db: db,
	}
}

func (r *TruckDB) CreateNewTruck(req entity.CreateNewTruckRequest) (entity.CreateNewTruckResponse, error) {
	var res entity.CreateNewTruckResponse
	q, err := r.db.Prepare(queryCreateNewTruck)
	if err != nil {
		return res, err
	}

	err = q.QueryRow(req.LicenseNumber, req.TruckType, req.PlateType, req.ProductionYear, true, false, req.STNK, req.KIR).Scan(&res.ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *TruckDB) GetTruckList(req entity.GetTruckListRequest) (entity.GetTruckListResponse, error) {
	var res entity.GetTruckListResponse
	query := queryGetAllTrucks

	if req.LicenseNumber != "" {
		req.LicenseNumber = "%" + req.LicenseNumber + "%"
	}
	if req.TruckType != "" {
		req.TruckType = "%" + req.TruckType + "%"
	}
	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Queryx(query, req.LicenseNumber, req.TruckType, offset, req.Limit)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := entity.Truck{}
		if err := rows.StructScan(&temp); err != nil {
			return res, err
		}
		res.Data.Rows = append(res.Data.Rows, temp)
	}

	return res, nil

}

func (r *TruckDB) GetTruckByID(req entity.GetTruckByIDRequest) (entity.GetTruckByIDResponse, error) {
	var res entity.GetTruckByIDResponse
	err := r.db.Get(&res, queryGetTruckById, req.ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *TruckDB) UpdateTruck(req entity.UpdateTruckRequest) (entity.UpdateTruckResponse, error) {
	_, err := r.db.Exec(queryUpdateTruck, req.LicenseNumber, req.TruckType, req.PlateType, req.ProductionYear, req.Status, req.STNK, req.KIR, req.ID)
	if err != nil {
		return entity.UpdateTruckResponse{}, err
	}
	return entity.UpdateTruckResponse{}, nil
}

func (r *TruckDB) DeleteTruck(req entity.DeleteTruckRequest) (entity.DeleteTruckResponse, error) {
	_, err := r.db.Exec(queryDeleteTruck, req.ID)
	if err != nil {
		return entity.DeleteTruckResponse{}, err
	}
	return entity.DeleteTruckResponse{}, nil
}
