package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) GetUser(c echo.Context) (err error) {

	return c.HTML(http.StatusOK, fmt.Sprintf("<BODY>Hello world - %s</BODY>", h.DB))

}
