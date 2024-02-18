package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/spleeroosh/go-translate/view/auth"
)

type AuthHandler struct{}

func (h AuthHandler) HandleAuthShow(c echo.Context) error {
	return render(c, auth.Show())
}
