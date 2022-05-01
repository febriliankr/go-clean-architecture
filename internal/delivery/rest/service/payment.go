package service

import (
	"net/http"
	"strconv"

	"github.com/febriliankr/go-clean-architecture/internal"
	"github.com/febriliankr/go-clean-architecture/internal/app"
	"github.com/febriliankr/go-clean-architecture/internal/entity"
	"github.com/labstack/echo/v4"
)

type PaymentService struct {
	uc internal.PaymentUC
}

func NewPaymentService(app *app.SeminarApp) *PaymentService {
	return &PaymentService{
		uc: app.Usecases.PaymentUC,
	}
}

func (s *PaymentService) CreateNewPaymentHandler(c echo.Context) error {
	if ok := validateUserGroup(c, entity.UserGroupTransporter); !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user token is not valid",
		})
	}
	var req entity.CreateNewPaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := s.uc.CreateNewPayment(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, res)
}

func (s *PaymentService) GetPaymentListHandler(c echo.Context) error {

	// if ok := validateUserGroup(c, entity.UserGroupTransporter); !ok {
	// 	return c.JSON(http.StatusUnauthorized, map[string]string{
	// 		"message": "user token is not valid",
	// 	})
	// }

	var req entity.GetPaymentListRequest

	req.Page, req.Limit = getPagination(c)

	res, err := s.uc.GetPaymentList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *PaymentService) GetPaymentByIDHandler(c echo.Context) error {
	if ok := validateUserGroup(c, entity.UserGroupTransporter); !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user token is not valid",
		})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	var req entity.GetPaymentByIDRequest
	req.ID = int64(id)

	res, err := s.uc.GetPaymentByID(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *PaymentService) UpdatePaymentHandler(c echo.Context) error {
	if ok := validateUserGroup(c, entity.UserGroupTransporter); !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user token is not valid",
		})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	var req entity.UpdatePaymentRequest
	req.ID = int64(id)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := s.uc.UpdatePayment(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *PaymentService) DeletePaymentHandler(c echo.Context) error {
	if ok := validateUserGroup(c, entity.UserGroupTransporter); !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user token is not valid",
		})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	var req entity.DeletePaymentRequest
	req.ID = int64(id)

	res, err := s.uc.DeletePayment(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
