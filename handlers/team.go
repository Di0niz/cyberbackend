package handlers

import (
	"fmt"
	"net/http"

	"github.com/Di0niz/cyberbackend/models"
	"github.com/labstack/echo"
)

func (h *Handler) GetTeam(c echo.Context) (err error) {

	return c.HTML(http.StatusOK, fmt.Sprintf("<BODY>GET /Team", h.DB))

}

func (h *Handler) PostTeam(c echo.Context) (err error) {

	t := new(models.Team)

	if err = c.Bind(t); err != nil {

		return c.JSON(http.StatusNotAcceptable, `{"err": "Request is not correct"}`)

	}

	err = h.DB.Create(t)

	if err != nil {
		return c.JSON(http.StatusNotAcceptable, `{"err": "Request is not correct"}`)
	}

	return c.JSON(http.StatusOK, t)

}

func (h *Handler) PutTeam(c echo.Context) (err error) {

	t := new(models.Team)

	if err = c.Bind(t); err != nil {

		err = h.DB.Update(t)

		if err == nil {
			return c.JSON(http.StatusOK, t)
		} else {
			return c.JSON(http.StatusNotAcceptable, err)
		}
	}
	return c.JSON(http.StatusNotAcceptable, err)

}

// получение списка комманд
func (h *Handler) GetListTeam(c echo.Context) (err error) {

	err = h.DB.List("team")

	return c.HTML(http.StatusOK, fmt.Sprintf("<BODY>Hello world - %s</BODY>", h.DB))

}

func (h *Handler) DeleteTeam(c echo.Context) (err error) {

	return c.HTML(http.StatusOK, fmt.Sprintf("<BODY>Hello world - %s</BODY>", h.DB))

}

func (h *Handler) GetListTeam(c echo.Context) (err error) {

	return c.HTML(http.StatusOK, fmt.Sprintf("<BODY>Hello world - %s</BODY>", h.DB))

}

func (h *Handler) GetTopTeam(c echo.Context) (err error) {

	return c.HTML(http.StatusOK, fmt.Sprintf("<BODY>Hello world - %s</BODY>", h.DB))

}
