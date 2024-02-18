package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/spleeroosh/go-translate/model"
	"github.com/spleeroosh/go-translate/view/translations"
)

type TranslationsHandler struct{}

func (h TranslationsHandler) HandleTranslationsShow(c echo.Context) error {
	return render(c, translations.Show(&model.Translation{
		Key:         "global.action.cancel",
		Lang:        "en",
		Translation: "Cancel",
	}))
}
