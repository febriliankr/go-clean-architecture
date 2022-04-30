package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/app"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type DriverService struct {
	uc internal.DriverUC
}

func NewDriverService(app *app.ShipmentApp) *DriverService {
	return &DriverService{
		uc: app.Usecases.DriverUC,
	}
}

func (s *DriverService) CreateNewDriverHandler(c echo.Context) error {
	if ok := validateUserGroup(c, entity.UserGroupTransporter); !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user token is not valid",
		})
	}
	var req entity.CreateNewDriverRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := s.uc.CreateNewDriver(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, res)
}

func (s *DriverService) GetDriverListHandler(c echo.Context) error {
	if ok := validateUserGroup(c, entity.UserGroupTransporter); !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user token is not valid",
		})
	}
	var req entity.GetDriverListRequest
	req.Name = c.QueryParam("name")
	req.Page, req.Limit = getPagination(c)

	res, err := s.uc.GetDriverList(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *DriverService) GetDriverByIDHandler(c echo.Context) error {
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
	var req entity.GetDriverByIDRequest
	req.ID = int64(id)

	res, err := s.uc.GetDriverByID(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *DriverService) UpdateDriverHandler(c echo.Context) error {
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
	var req entity.UpdateDriverRequest
	req.ID = int64(id)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := s.uc.UpdateDriver(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *DriverService) DeleteDriverHandler(c echo.Context) error {
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
	var req entity.DeleteDriverRequest
	req.ID = int64(id)

	res, err := s.uc.DeleteDriver(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
