package controllers

import (
	// "github.com/labstack/echo"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/vshaveyko/h24-static-api/models"
	// "json"
)

type StaticTranslationsController struct {
}

func (u StaticTranslationsController) Index(c echo.Context) interface{} {

	result := make(map[string]string)

	locale := c.QueryParam("lang")

	if locale == "" {
		locale = "ru"
	}

	var static_translations []models.StaticTranslation

	models.DB.Joins(`
			INNER JOIN static_translation_translations
			   ON static_translation_translations.static_translation_id = static_translations.id
				AND static_translation_translations.locale = ?
		`, locale).Select("DISTINCT ON (text_alias) text_alias, static_translation_translations.static_translation as static_translation").Find(&static_translations)

	for _, tl := range static_translations {
		fmt.Print("Alias", tl.TextAlias, "TRANSLATION", tl.StaticTranslation, "\n")
		result[tl.TextAlias] = tl.StaticTranslation
	}

	return result

}

func (u StaticTranslationsController) Create(c echo.Context) (interface{}, error) {
	return nil, errors.New("Not implememneted")
}
func (u StaticTranslationsController) Show(c echo.Context) (interface{}, error) {
	return nil, errors.New("Not implememneted")
}
func (u StaticTranslationsController) Destroy(c echo.Context) error {
	return errors.New("Not implememneted")
}
func (u StaticTranslationsController) Update(c echo.Context) (interface{}, error) {
	return nil, errors.New("Not implememneted")
}
