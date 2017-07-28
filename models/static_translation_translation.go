package models

type StaticTranslationTranslation struct {
	Base

	Locale              string
	StaticTranslation   string
	Record              StaticTranslation
	StaticTranslationId int
}
