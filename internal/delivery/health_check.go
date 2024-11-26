package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (s Server) healthCheck(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good! " + time.Now().Format("2006-01-02 15:04:05"),
	})
}
