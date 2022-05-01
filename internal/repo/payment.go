package repo

import (
	"github.com/febriliankr/go-clean-architecture/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PaymentDB struct {
	db *sqlx.DB
}

func NewPaymentDB(db *sqlx.DB) *PaymentDB {
	return &PaymentDB{
		db: db,
	}
}

func (r *PaymentDB) CreateNewPayment(req entity.CreateNewPaymentRequest) (entity.CreateNewPaymentResponse, error) {
	var res entity.CreateNewPaymentResponse
	q, err := r.db.Prepare(queryCreateNewPayment)
	if err != nil {
		return res, err
	}

	err = q.QueryRow().Scan(&res.ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *PaymentDB) GetPaymentList(req entity.GetPaymentListRequest) (entity.GetPaymentListResponse, error) {
	var res entity.GetPaymentListResponse
	query := queryGetAllPayments

	offset := (req.Page - 1) * req.Limit

	rows, err := r.db.Queryx(query, offset, req.Limit)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {

		temp := entity.Payment{}

		if err := rows.StructScan(&temp); err != nil {
			return res, err
		}

		res.Data.Rows = append(res.Data.Rows, temp)
	}

	return res, nil

}

func (r *PaymentDB) GetPaymentByID(req entity.GetPaymentByIDRequest) (entity.GetPaymentByIDResponse, error) {
	var res entity.GetPaymentByIDResponse
	err := r.db.Get(&res, queryGetPaymentById, req.ID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *PaymentDB) UpdatePayment(req entity.UpdatePaymentRequest) (entity.UpdatePaymentResponse, error) {
	_, err := r.db.Exec(queryUpdatePayment, req.PaymentExpiredTimestamp, req.PaymentReminderTimestamp, req.EventReminderTimestamp, req.BuktiPembayaranUrl, req.ID)

	if err != nil {
		return entity.UpdatePaymentResponse{}, err
	}
	return entity.UpdatePaymentResponse{}, nil
}

func (r *PaymentDB) DeletePayment(req entity.DeletePaymentRequest) (entity.DeletePaymentResponse, error) {
	_, err := r.db.Exec(queryDeletePayment, req.ID)
	if err != nil {
		return entity.DeletePaymentResponse{}, err
	}
	return entity.DeletePaymentResponse{}, nil
}
