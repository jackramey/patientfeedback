package service

import (
	"net/http"

	"patientfeedback/storage"

	"github.com/labstack/echo/v4"
)

type adminHandler struct {
	store storage.ResourceStore
}

func (handler adminHandler) DumpDB(c echo.Context) error {
	dbDump := handler.store.DumpDB()
	return c.JSON(http.StatusOK, dbDump)
}
