package models

type StaticTranslation struct {
	Base

	TextAlias         string                         `gorm:"type:text"`
	StaticTranslation string                         `gorm:"type:text"`
	Translations      []StaticTranslationTranslation `gorm:"ForeignKey:static_translation_id"`
}
