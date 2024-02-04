package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/spleeroosh/go-translate/view/home"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	return render(c, home.Show())
}
