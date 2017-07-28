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

func allLocalesResult() map[string]map[string]string {

	result := map[string]map[string]string{}

	var static_translations []models.StaticTranslation

	models.DB.Preload("Translations").Find(&static_translations)

	for _, tl := range static_translations {

		for _, locale_tl := range tl.Translations {
			if _, ok := result[locale_tl.Locale]; !ok {
				result[locale_tl.Locale] = map[string]string{}
			}

			result[locale_tl.Locale][tl.TextAlias] = locale_tl.StaticTranslation

		}

	}

	return result

}

func localeResult(locale string) map[string]string {

	result := make(map[string]string)

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

func (u StaticTranslationsController) Index(c echo.Context) interface{} {

	locale := c.QueryParam("lang")

	if locale == "" {
		return allLocalesResult()
	} else {
		return localeResult(locale)
	}

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
