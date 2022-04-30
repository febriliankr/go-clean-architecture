package service

import (
	"github.com/ridwanakf/tms-backend/internal/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/app"
)

type ShipmentService struct {
	uc internal.ShipmentUC
}

func NewShipmentService(app *app.ShipmentApp) *ShipmentService {
	return &ShipmentService{
		uc: app.Usecases.ShipmentUC,
	}
}

func (s *ShipmentService) CreateNewShipmentHandler(c echo.Context) error {
	if ok := validateUserGroup(c, entity.UserGroupShipper); !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user token is not valid",
		})
	}
	var req entity.CreateNewShipmentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := s.uc.CreateNewShipment(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, res)
}

func (s *ShipmentService) AllocateShipmentHandler(c echo.Context) error {
	if ok := validateUserGroup(c, entity.UserGroupShipper); !ok {
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
	var req entity.AllocateShipmentRequest
	req.ID = int64(id)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := s.uc.AllocateShipment(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, res)
}

func (s *ShipmentService) GetShipmentListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *ShipmentService) GetShipmentByIDHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *ShipmentService) UpdateShipmentHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *ShipmentService) DeleteShipmentHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
