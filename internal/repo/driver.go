package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type DriverDB struct {
	db *sqlx.DB
}

func NewDriverDB(db *sqlx.DB) *DriverDB {
	return &DriverDB{
		db: db,
	}
}

func (r *DriverDB) CreateNewDriver(req entity.CreateNewDriverRequest) (entity.CreateNewDriverResponse, error) {
	var res entity.CreateNewDriverResponse
	q, err := r.db.Prepare(queryCreateNewDriver)
	if err != nil {
		return res, err
	}

	err = q.QueryRow(req.Name, req.Phone, req.CardID, req.DriverLicense, true, false).Scan(&res.ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *DriverDB) GetDriverList(req entity.GetDriverListRequest) (entity.GetDriverListResponse, error) {
	var res entity.GetDriverListResponse
	query := queryGetAllDrivers

	if req.Name != "" {
		req.Name = "%" + req.Name + "%"
	}

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Queryx(query, req.Name, offset, req.Limit)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := entity.Driver{}
		if err := rows.StructScan(&temp); err != nil {
			return res, err
		}
		res.Data.Rows = append(res.Data.Rows, temp)
	}

	return res, nil
}

func (r *DriverDB) GetDriverByID(req entity.GetDriverByIDRequest) (entity.GetDriverByIDResponse, error) {
	var res entity.GetDriverByIDResponse
	err := r.db.Get(&res, queryGetDriverById, req.ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *DriverDB) UpdateDriver(req entity.UpdateDriverRequest) (entity.UpdateDriverResponse, error) {
	_, err := r.db.Exec(queryUpdateDriver, req.Name, req.Phone, req.DriverLicense, req.Status, req.Status, req.ShipmentStatus, req.ID)
	if err != nil {
		return entity.UpdateDriverResponse{}, err
	}
	return entity.UpdateDriverResponse{}, nil
}

func (r *DriverDB) DeleteDriver(req entity.DeleteDriverRequest) (entity.DeleteDriverResponse, error) {
	_, err := r.db.Exec(queryDeleteDriver, req.ID)
	if err != nil {
		return entity.DeleteDriverResponse{}, err
	}
	return entity.DeleteDriverResponse{}, nil
}
