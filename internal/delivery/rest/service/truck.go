package service

import (
	"net/http"
	"strconv"

	"github.com/febriliankr/go-clean-architecture/internal"
	"github.com/febriliankr/go-clean-architecture/internal/app"
	"github.com/febriliankr/go-clean-architecture/internal/entity"
	"github.com/labstack/echo/v4"
)

type TruckService struct {
	uc internal.TruckUC
}

func NewTruckService(app *app.ShipmentApp) *TruckService {
	return &TruckService{
		uc: app.Usecases.TruckUC,
	}
}

func (s *TruckService) CreateNewTruckHandler(c echo.Context) error {
	if ok := validateUserGroup(c, entity.UserGroupTransporter); !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user token is not valid",
		})
	}
	var req entity.CreateNewTruckRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := s.uc.CreateNewTruck(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, res)
}

func (s *TruckService) GetTruckListHandler(c echo.Context) error {

	if ok := validateUserGroup(c, entity.UserGroupTransporter); !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user token is not valid",
		})
	}
	var req entity.GetTruckListRequest
	req.LicenseNumber = c.QueryParam("license")
	req.TruckType = c.QueryParam("type")
	req.Page, req.Limit = getPagination(c)

	res, err := s.uc.GetTruckList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *TruckService) GetTruckByIDHandler(c echo.Context) error {
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
	var req entity.GetTruckByIDRequest
	req.ID = int64(id)

	res, err := s.uc.GetTruckByID(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *TruckService) UpdateTruckHandler(c echo.Context) error {
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
	var req entity.UpdateTruckRequest
	req.ID = int64(id)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := s.uc.UpdateTruck(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *TruckService) DeleteTruckHandler(c echo.Context) error {
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
	var req entity.DeleteTruckRequest
	req.ID = int64(id)

	res, err := s.uc.DeleteTruck(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
