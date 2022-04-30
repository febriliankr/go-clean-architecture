package service

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func getBearerToken(c echo.Context) string {
	reqToken := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	return reqToken
}

func validateUserGroup(c echo.Context, group string) bool {
	token := getBearerToken(c)
	return token == group
}

func getPagination(c echo.Context) (int64, int64) {
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")
	page := 1
	limit := 10
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}
	return int64(page), int64(limit)
}
