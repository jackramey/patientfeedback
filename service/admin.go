package service

import (
	"net/http"

	"patientfeedback/storage"

	"github.com/labstack/echo/v4"
)

type securedHandler struct {
	store storage.ResourceStore
}

func (handler securedHandler) DumpDB(c echo.Context) error {
	dbDump := handler.store.DumpDB()
	return c.JSON(http.StatusOK, dbDump)
}

func (handler securedHandler) ProcessBundle(c echo.Context) error {
	req := c.Request()
	return storage.LoadData(req.Body, handler.store)
}
